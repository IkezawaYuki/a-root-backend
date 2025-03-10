package usecase

import (
	"IkezawaYuki/a-root-backend/domain/arootErr"
	"IkezawaYuki/a-root-backend/domain/entity"
	"IkezawaYuki/a-root-backend/domain/model"
	"IkezawaYuki/a-root-backend/interface/dto/external"
	"IkezawaYuki/a-root-backend/interface/dto/req"
	"IkezawaYuki/a-root-backend/interface/dto/res"
	"IkezawaYuki/a-root-backend/interface/filter"
	"IkezawaYuki/a-root-backend/interface/repository"
	"IkezawaYuki/a-root-backend/service"
	"context"
	"golang.org/x/crypto/bcrypt"
	"slices"
)

type CustomerUsecase interface {
	FetchAndPost(ctx context.Context, customerID int) (*res.Message, error)
	Login(ctx context.Context, user req.User) (*res.Auth, error)
	GetCustomer(ctx context.Context, customerID int) (*res.Customer, error)
	GetPosts(ctx context.Context, customerID int, req req.PostQuery) (*res.Posts, error)
	FetchInstagramPosts(ctx context.Context, customerID int) (*external.InstagramPosts, error)
	GenerateLongToken(ctx context.Context, customerID int, body req.AuthTokenBody) (*res.Message, error)
	RefreshToken(ctx context.Context, customerID int) (*res.Message, error)
	Register(ctx context.Context, customerID int, body req.RegisterCustomerBody) (*res.Customer, error)
}

type customerUsecase struct {
	baseRepo        repository.BaseRepository
	postRepo        repository.PostRepository
	customerRepo    repository.CustomerRepository
	rodutRepo       repository.RodutRepository
	customerService service.CustomerService
	authService     service.AuthService
	postService     service.PostService
	graphApi        service.GraphAPI
	fileTransfer    service.FileService
	slackService    service.SlackService
	openaiService   service.OpenaiService
}

func NewCustomerUsecase(
	baseRepo repository.BaseRepository,
	postRepo repository.PostRepository,
	customerRepo repository.CustomerRepository,
	customerSrv service.CustomerService,
	authSrv service.AuthService,
	postService service.PostService,
	graphApi service.GraphAPI,
	fileTransfer service.FileService,
	rodutRepo repository.RodutRepository,
	slackService service.SlackService,
	openaiService service.OpenaiService,
) CustomerUsecase {
	return &customerUsecase{
		baseRepo:        baseRepo,
		postRepo:        postRepo,
		customerRepo:    customerRepo,
		customerService: customerSrv,
		authService:     authSrv,
		postService:     postService,
		graphApi:        graphApi,
		fileTransfer:    fileTransfer,
		rodutRepo:       rodutRepo,
		slackService:    slackService,
		openaiService:   openaiService,
	}
}

func (c *customerUsecase) FindAll(ctx context.Context, query req.CustomerQuery) (*res.Customers, error) {

	var instagramTokenStatus *entity.InstagramTokenStatus
	if query.InstagramTokenStatus != nil {
		i := entity.ConvertToInstagramTokenStatus(query.InstagramTokenStatus)
		if i == entity.InstagramTokenStatusUnknown {
			return nil, arootErr.ErrInvalidInstagramTokenStatus
		}
		instagramTokenStatus = &i
	}

	var paymentStatus *entity.PaymentStatus
	if query.PaymentType != nil {
		p := entity.ConvertToPaymentStatus(query.PaymentType)
		if p == entity.PaymentStatusUnknown {
			return nil, arootErr.ErrInvalidInstagramTokenStatus
		}
		paymentStatus = &p
	}

	var paymentType *entity.PaymentType
	if query.PaymentType != nil {
		p := entity.ConvertToPaymentType(query.PaymentType)
		if p == entity.PaymentTypeUnknown {
			return nil, arootErr.ErrInvalidInstagramTokenStatus
		}
		paymentType = &p
	}

	f := &filter.CustomerFilter{
		InstagramTokenStatus: instagramTokenStatus,
		Email:                query.Email,
		PartialName:          query.PartialName,
		PaymentType:          paymentType,
		PaymentStatus:        paymentStatus,
		Limit:                query.Limit,
		Offset:               query.Offset,
	}
	customers, err := c.customerRepo.Get(ctx, f)
	if err != nil {
		return nil, err
	}
	count, err := c.customerRepo.Count(ctx, f)
	if err != nil {
		return nil, err
	}
	return res.GetCustomers(customers, count), nil
}

func (c *customerUsecase) GetCustomer(ctx context.Context, id int) (*res.Customer, error) {
	customer, err := c.customerRepo.First(ctx, &filter.CustomerFilter{
		ID: &id,
	})
	if err != nil {
		return nil, err
	}
	return res.GetCustomer(customer), nil
}

func (c *customerUsecase) Login(ctx context.Context, user req.User) (*res.Auth, error) {
	customer, err := c.customerService.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, arootErr.ErrAuthorization
	}
	if err := c.authService.CheckPassword(user, customer.Password); err != nil {
		return nil, arootErr.ErrAuthorization
	}
	token, err := c.authService.GenerateJWTCustomer(customer)
	if err != nil {
		return nil, err
	}
	return &res.Auth{
		Token:  token,
		UserID: int(customer.ID),
	}, nil
}

func (c *customerUsecase) FetchAndPost(ctx context.Context, customerID int) (*res.Message, error) {
	customer, err := c.customerService.FindByID(ctx, customerID)
	if err != nil {
		return nil, arootErr.ErrNotFound
	}

	if !customer.IsLinked() {
		return nil, arootErr.ErrInstagramNotLinked
	}

	// インスタグラムの投稿を最新から50件取得する
	instagramPosts, err := c.graphApi.GetInstagramPosts(ctx, *customer.FacebookToken, *customer.InstagramAccountID)
	if err != nil {
		return nil, err
	}

	// 投稿済みのMediaIDを一覧で取得する
	linkedMediaIDs, err := c.postService.GetLinkedMediaIDs(ctx, customerID)
	if err != nil {
		return nil, err
	}

	for _, instagramMedia := range instagramPosts.Media.Data {
		if slices.Contains(linkedMediaIDs, instagramMedia.ID) {
			// 連携済みのためスキップ
			continue
		}

		// 一時フォルダを作る
		if err := c.fileTransfer.MakeTempDirectory(customerID); err != nil {
			return nil, err
		}

		// メディアをダウンロード
		fileList, err := c.fileTransfer.DownloadMediaFiles(ctx, customerID, instagramMedia)
		if err != nil {
			return nil, err
		}

		// ワードプレスにメディアをアップロード
		wordpressMedia, err := c.rodutRepo.UploadMedias(ctx, customer.WordpressUrl, fileList)
		if err != nil {
			return nil, err
		}

		// ワードプレスに記事を投稿
		createPost := external.NewWordpressPost(instagramMedia, wordpressMedia)
		wordpressResp, err := c.rodutRepo.CreatePost(ctx, customer.WordpressUrl, createPost)
		if err != nil {
			return nil, err
		}
		post := &model.Post{
			CustomerID:       customerID,
			InstagramMediaID: instagramMedia.ID,
			InstagramLink:    instagramMedia.MediaURL,
			WordpressMediaID: wordpressResp.PostId,
			WordpressLink:    wordpressResp.PostUrl,
		}

		// 投稿をDBに保存
		if err := c.postRepo.Save(ctx, post); err != nil {
			return nil, err
		}

		// 一時ディレクトリを削除
		if err := c.fileTransfer.RemoveTempDirectory(customerID); err != nil {
			return nil, err
		}

		_ = c.slackService.Log(ctx, customer, post)
	}
	return &res.Message{Message: "ok"}, nil
}

func (c *customerUsecase) GetPosts(ctx context.Context, customerID int, req req.PostQuery) (*res.Posts, error) {
	f := &filter.PostFilter{
		CustomerID: &customerID,
		Limit:      req.Limit,
		Offset:     req.Offset,
	}
	posts, err := c.postRepo.Get(ctx, f)
	if err != nil {
		return nil, err
	}
	counts, err := c.postRepo.Count(ctx, f)
	if err != nil {
		return nil, err
	}
	return res.GetPosts(posts, counts), nil
}

func (c *customerUsecase) FetchInstagramPosts(ctx context.Context, customerID int) (*external.InstagramPosts, error) {
	customer, err := c.customerService.FindByID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if !customer.IsLinked() {
		return nil, arootErr.ErrInstagramNotLinked
	}
	posts, err := c.graphApi.GetInstagramPosts(ctx, *customer.FacebookToken, *customer.InstagramAccountID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (c *customerUsecase) GenerateLongToken(ctx context.Context, customerID int, body req.AuthTokenBody) (*res.Message, error) {
	customer, err := c.customerService.FindByID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	authTokenResp, err := c.graphApi.GetOAuthAccessToken(ctx, body.AccessToken)
	if err != nil {
		return nil, err
	}
	customer.InstagramTokenStatus = entity.InstagramTokenStatusActive
	customer.FacebookToken = &authTokenResp.AccessToken
	if err := c.customerRepo.Save(ctx, customer); err != nil {
		return nil, err
	}
	return &res.Message{Message: "ok"}, nil
}

func (c *customerUsecase) RefreshToken(ctx context.Context, customerID int) (resp *res.Message, err error) {
	tx := c.baseRepo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	customer, err := c.customerRepo.FirstTx(ctx, &filter.CustomerFilter{ID: &customerID}, tx)
	if err != nil {
		return nil, err
	}
	if !customer.IsLinked() {
		return nil, arootErr.ErrInstagramNotLinked
	}
	authTokenResp, err := c.graphApi.GetOAuthAccessToken(ctx, *customer.FacebookToken)
	if err != nil {
		return nil, err
	}
	customer.FacebookToken = &authTokenResp.AccessToken
	if err := c.customerRepo.Save(ctx, customer); err != nil {
		return nil, err
	}

	return &res.Message{Message: "ok"}, nil
}

func (c *customerUsecase) Register(ctx context.Context, customerID int, body req.RegisterCustomerBody) (*res.Customer, error) {
	customer, err := c.customerService.FindByID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	titleResp, err := c.rodutRepo.GetTitle(ctx, body.WordpressURL)
	if err != nil {
		return nil, err
	}
	customer.Name = titleResp.Title
	customer.WordpressUrl = body.WordpressURL
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	customer.Password = string(passwordHash)
	if err := c.customerRepo.Save(ctx, customer); err != nil {
		return nil, err
	}
	return res.GetCustomer(customer), nil
}

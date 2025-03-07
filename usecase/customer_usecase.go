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
	"IkezawaYuki/a-root-backend/util"
	"context"
	"fmt"
	"slices"
)

type CustomerUsecase interface {
	FetchAndPost(ctx context.Context, customerID int) (*res.Message, error)
	Login(ctx context.Context, user req.User) (*res.Auth, error)
	GetCustomer(ctx context.Context, customerID int) (*res.Customer, error)
	GetPosts(ctx context.Context, customerID int, req req.PostQuery) (*res.Posts, error)
	FetchInstagramPosts(ctx context.Context, customerID int) (*external.InstagramPosts, error)
	GenerateLongToken(ctx context.Context, customerID int, body req.AuthTokenBody) (*res.Message, error)
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
	}
}

func (c *customerUsecase) FindAll(ctx context.Context, query req.CustomerQuery) (*res.Customers, error) {

	var instagramTokenStatus *model.InstagramTokenStatus
	if query.InstagramTokenStatus != nil {
		i := model.ConvertToInstagramTokenStatus(query.InstagramTokenStatus)
		if i == model.InstagramTokenStatusUnknown {
			return nil, arootErr.ErrInvalidInstagramTokenStatus
		}
		instagramTokenStatus = &i
	}

	var paymentStatus *model.PaymentStatus
	if query.PaymentType != nil {
		p := model.ConvertToPaymentStatus(query.PaymentType)
		if p == model.PaymentStatusUnknown {
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
		return nil, err
	}
	if err := c.authService.CheckPassword(user, customer.Password); err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}
	token, err := c.authService.GenerateJWTCustomer(customer)
	if err != nil {
		return nil, err
	}
	return &res.Auth{Token: token}, nil
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

	linkedMediaIDs, err := c.postService.GetLinkedMediaIDs(ctx, customerID)
	if err != nil {
		return nil, err
	}

	for _, instagramMedia := range instagramPosts.Media.Data {
		if slices.Contains(linkedMediaIDs, instagramMedia.ID) {
			// 連携済みのためスキップ
			continue
		}

		// 一時フォルダを作り、メディアをダウンロード
		if err := c.fileTransfer.MakeTempDirectory(customerID); err != nil {
			return nil, err
		}

		fileList, err := c.fileTransfer.DownloadMediaFiles(ctx, customerID, instagramMedia)
		if err != nil {
			return nil, err
		}

		// ワードプレスにメディアをアップロード
		wordpressMedia, err := c.rodutRepo.UploadMedias(ctx, customer.WordpressUrl, fileList)
		if err != nil {
			return nil, err
		}

		createPost := external.NewWordpressPost(instagramMedia, wordpressMedia)
		wordpressResp, err := c.rodutRepo.CreatePost(ctx, customer.WordpressUrl, createPost)
		if err != nil {
			return nil, err
		}

		if err := c.postRepo.Save(ctx, &model.Post{
			CustomerID:       customerID,
			InstagramMediaID: instagramMedia.ID,
			InstagramLink:    instagramMedia.MediaURL,
			WordpressMediaID: wordpressResp.PostId,
			WordpressLink:    wordpressResp.PostUrl,
		}); err != nil {
			return nil, err
		}

		if err := c.fileTransfer.RemoveTempDirectory(customerID); err != nil {
			return nil, err
		}
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
	customer.InstagramTokenStatus = util.Pointer(model.InstagramTokenStatusActive)
	customer.FacebookToken = &authTokenResp.AccessToken
	if err := c.customerRepo.Save(ctx, customer); err != nil {
		return nil, err
	}
	return &res.Message{Message: "ok"}, nil
}

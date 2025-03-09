package usecase

import (
	"IkezawaYuki/a-root-backend/domain/arootErr"
	"IkezawaYuki/a-root-backend/domain/model"
	"IkezawaYuki/a-root-backend/interface/dto/req"
	"IkezawaYuki/a-root-backend/interface/dto/res"
	"IkezawaYuki/a-root-backend/interface/filter"
	"IkezawaYuki/a-root-backend/interface/repository"
	"IkezawaYuki/a-root-backend/service"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	baseRepository  repository.BaseRepository
	adminRepo       repository.AdminRepository
	customerRepo    repository.CustomerRepository
	postRepo        repository.PostRepository
	adminService    service.AdminService
	authService     service.AuthService
	customerService service.CustomerService
	postService     service.PostService
}

type AdminUsecase interface {
	Login(ctx context.Context, user req.User) (*res.Auth, error)
	CreateCustomer(ctx context.Context, body req.CreateCustomerBody) (resp *res.Customer, err error)
	CreateAdmin(ctx context.Context, body req.CreateAdminBody) (*res.Admin, error)
	GetCustomers(ctx context.Context, query req.CustomerQuery) (*res.Customers, error)
	GetCustomer(ctx context.Context, customerID int) (*res.Customer, error)
	GetAdmin(ctx context.Context, id int) (*res.Admin, error)
	GetAdmins(ctx context.Context, query req.AdminQuery) (*res.Admins, error)
	GetPosts(ctx context.Context, customerID int, query req.PostQuery) (*res.Posts, error)
	DeleteCustomer(ctx context.Context, customerID int) (*res.Message, error)
	DeleteAdmin(ctx context.Context, adminId int) (*res.Message, error)
	UpdateCustomer(ctx context.Context, customerID int, body req.UpdateCustomerBody) (*res.Customer, error)
	UpdateAdmin(ctx context.Context, adminID int, body req.UpdateAdminBody) (*res.Admin, error)
}

func NewAdminUsecase(
	baseRepo repository.BaseRepository,
	adminRepo repository.AdminRepository,
	customerRepo repository.CustomerRepository,
	postRepo repository.PostRepository,
	adminSrv service.AdminService,
	authSrv service.AuthService,
	customerService service.CustomerService,
) AdminUsecase {
	return &adminUsecase{
		baseRepository:  baseRepo,
		adminRepo:       adminRepo,
		customerRepo:    customerRepo,
		postRepo:        postRepo,
		adminService:    adminSrv,
		authService:     authSrv,
		customerService: customerService,
	}
}

func (a *adminUsecase) CreateCustomer(ctx context.Context, body req.CreateCustomerBody) (resp *res.Customer, err error) {
	tx := a.baseRepository.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	used, err := a.customerService.IsUsedEmailAddress(ctx, body.Email, tx)
	if err != nil {
		return nil, err
	}
	if used {
		return nil, arootErr.ErrEmailUsed
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	customer := &model.Customer{
		Name:         body.Name,
		Email:        body.Email,
		Password:     string(passwordHash),
		WordpressUrl: body.WordpressURL,
	}
	err = a.customerRepo.SaveTx(ctx, customer, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return res.GetCustomer(customer), nil
}

func (a *adminUsecase) CreateAdmin(ctx context.Context, body req.CreateAdminBody) (*res.Admin, error) {
	used, err := a.adminService.IsUsedEmailAddress(ctx, body.Email)
	if err != nil {
		return nil, err
	}
	if used {
		return nil, arootErr.ErrEmailUsed
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin := &model.Admin{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(passwordHash),
	}
	err = a.adminRepo.Save(ctx, admin)
	if err != nil {
		return nil, err
	}
	return res.GetAdmin(admin), nil
}

func (a *adminUsecase) Login(ctx context.Context, user req.User) (*res.Auth, error) {
	customer, err := a.adminService.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, arootErr.ErrNotFound
	}
	if err := a.authService.CheckPassword(user, customer.Password); err != nil {
		return nil, err
	}
	token, err := a.authService.GenerateJWTAdmin(customer)
	if err != nil {
		return nil, err
	}
	return &res.Auth{Token: token}, nil
}

func (a *adminUsecase) GetCustomers(ctx context.Context, query req.CustomerQuery) (*res.Customers, error) {
	f := &filter.CustomerFilter{
		Email:       query.Email,
		PartialName: query.PartialName,
		Limit:       query.Limit,
		Offset:      query.Offset,
	}
	customers, err := a.customerRepo.Get(ctx, f)
	if err != nil {
		return nil, err
	}
	count, err := a.customerRepo.Count(ctx, f)
	if err != nil {
		return nil, err
	}
	return res.GetCustomers(customers, count), nil
}

func (a *adminUsecase) GetCustomer(ctx context.Context, id int) (*res.Customer, error) {
	customer, err := a.customerService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.GetCustomer(customer), nil
}

func (a *adminUsecase) GetAdmins(ctx context.Context, query req.AdminQuery) (*res.Admins, error) {
	f := &filter.AdminFilter{
		Email:       query.Email,
		PartialName: query.PartialName,
		Limit:       query.Limit,
		Offset:      query.Offset,
	}
	admins, err := a.adminRepo.Get(ctx, f)
	if err != nil {
		return nil, err
	}
	count, err := a.adminRepo.Count(ctx, f)
	if err != nil {
		return nil, err
	}
	return res.GetAdmins(admins, count), nil
}

func (a *adminUsecase) GetAdmin(ctx context.Context, id int) (*res.Admin, error) {
	admin, err := a.adminService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.GetAdmin(admin), nil
}

func (a *adminUsecase) GetPosts(ctx context.Context, customerID int, query req.PostQuery) (*res.Posts, error) {
	f := &filter.PostFilter{
		CustomerID: &customerID,
		Limit:      query.Limit,
		Offset:     query.Offset,
	}
	posts, err := a.postRepo.Get(ctx, f)
	if err != nil {
		return nil, err
	}
	count, err := a.postRepo.Count(ctx, f)
	if err != nil {
		return nil, err
	}
	return res.GetPosts(posts, count), nil
}

func (a *adminUsecase) DeleteCustomer(ctx context.Context, customerID int) (*res.Message, error) {
	_, err := a.customerService.FindByID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	err = a.customerRepo.Delete(ctx, &filter.CustomerFilter{
		ID: &customerID,
	})
	if err != nil {
		return nil, err
	}
	return &res.Message{Message: "ok"}, nil
}

func (a *adminUsecase) DeleteAdmin(ctx context.Context, adminId int) (*res.Message, error) {
	_, err := a.adminService.FindByID(ctx, adminId)
	if err != nil {
		return nil, err
	}
	err = a.adminRepo.Delete(ctx, &filter.AdminFilter{
		ID: &adminId,
	})
	if err != nil {
		return nil, err
	}
	return &res.Message{Message: "ok"}, nil
}

func (a *adminUsecase) UpdateCustomer(ctx context.Context, customerID int, body req.UpdateCustomerBody) (*res.Customer, error) {
	customer, err := a.customerService.FindByID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	customer.Name = body.Name
	err = a.customerRepo.Save(ctx, customer)
	if err != nil {
		return nil, err
	}
	return res.GetCustomer(customer), nil
}

func (a *adminUsecase) UpdateAdmin(ctx context.Context, adminID int, body req.UpdateAdminBody) (*res.Admin, error) {
	admin, err := a.adminService.FindByID(ctx, adminID)
	if err != nil {
		return nil, err
	}
	admin.Name = body.Name
	err = a.adminRepo.Save(ctx, admin)
	if err != nil {
		return nil, err
	}
	return res.GetAdmin(admin), nil
}

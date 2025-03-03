package service

import (
	"IkezawaYuki/a-root-backend/domain/arootErr"
	"IkezawaYuki/a-root-backend/domain/model"
	"IkezawaYuki/a-root-backend/interface/filter"
	"IkezawaYuki/a-root-backend/interface/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type adminService struct {
	customerRepository repository.CustomerRepository
	adminRepository    repository.AdminRepository
}

type AdminService interface {
	FindByID(ctx context.Context, id int) (*model.Admin, error)
	FindByEmail(ctx context.Context, email string) (*model.Admin, error)
	FindAll(ctx context.Context) ([]*model.Admin, error)
	IsUsedEmailAddress(ctx context.Context, email string) (bool, error)
}

func NewAdminService(customerRepo repository.CustomerRepository, adminRepo repository.AdminRepository) AdminService {
	return &adminService{
		customerRepository: customerRepo,
		adminRepository:    adminRepo,
	}
}

func (a *adminService) FindAll(ctx context.Context) ([]*model.Admin, error) {
	return a.adminRepository.Get(ctx, nil)
}

func (a *adminService) FindByEmail(ctx context.Context, email string) (*model.Admin, error) {
	return a.adminRepository.First(ctx, &filter.AdminFilter{Email: &email})
}

func (a *adminService) FindByID(ctx context.Context, id int) (*model.Admin, error) {
	return a.adminRepository.First(ctx, &filter.AdminFilter{ID: &id})
}

func (a *adminService) CreateAdmin(ctx context.Context, admin *model.Admin) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	adminModel := model.Admin{
		Name:     admin.Name,
		Email:    admin.Email,
		Password: string(passwordHash),
	}
	if err := a.adminRepository.Save(ctx, &adminModel); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return arootErr.ErrDuplicateEmail
		}
		return err
	}
	return nil
}

func (a *adminService) IsUsedEmailAddress(ctx context.Context, email string) (bool, error) {
	admins, err := a.adminRepository.Get(ctx, &filter.AdminFilter{Email: &email})
	if err != nil {
		return false, err
	}
	return len(admins) > 0, nil
}

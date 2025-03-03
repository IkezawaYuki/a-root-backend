package repository

import (
	"IkezawaYuki/a-root-backend/domain/model"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/interface/filter"
	"context"
)

type AdminRepository interface {
	Get(ctx context.Context, f *filter.AdminFilter) ([]*model.Admin, error)
	First(ctx context.Context, f *filter.AdminFilter) (*model.Admin, error)
	GetTx(ctx context.Context, f *filter.AdminFilter, tx infrastructure.Transaction) ([]*model.Admin, error)
	FirstTx(ctx context.Context, f *filter.AdminFilter, tx infrastructure.Transaction) (*model.Admin, error)
	Save(ctx context.Context, admin *model.Admin) error
	Count(ctx context.Context, f *filter.AdminFilter) (int, error)
	Delete(ctx context.Context, f *filter.AdminFilter) error
}

type adminRepository struct {
	dbDriver infrastructure.DBDriver
}

func NewAdminRepository(dbDriver infrastructure.DBDriver) AdminRepository {
	return &adminRepository{dbDriver: dbDriver}
}

func (a *adminRepository) Get(ctx context.Context, f *filter.AdminFilter) ([]*model.Admin, error) {
	var admins []*model.Admin
	err := a.dbDriver.Find(ctx, &admins, f)
	if err != nil {
		return nil, err
	}
	return admins, err
}

func (a *adminRepository) First(ctx context.Context, f *filter.AdminFilter) (*model.Admin, error) {
	var admin model.Admin
	err := a.dbDriver.First(ctx, &admin, f)
	return &admin, err
}

func (a *adminRepository) GetTx(ctx context.Context, f *filter.AdminFilter, tx infrastructure.Transaction) ([]*model.Admin, error) {
	var admins []*model.Admin
	err := a.dbDriver.FindTx(ctx, &admins, f, tx)
	if err != nil {
		return nil, err
	}
	return admins, err
}

func (a *adminRepository) FirstTx(ctx context.Context, f *filter.AdminFilter, tx infrastructure.Transaction) (*model.Admin, error) {
	var admin model.Admin
	err := a.dbDriver.FirstTx(ctx, &admin, f, tx)
	return &admin, err
}

func (a *adminRepository) Save(ctx context.Context, admin *model.Admin) error {
	return a.dbDriver.Save(ctx, admin)
}

func (a *adminRepository) Count(ctx context.Context, f *filter.AdminFilter) (int, error) {
	count, err := a.dbDriver.Count(ctx, &model.Admin{}, f)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (a *adminRepository) Delete(ctx context.Context, f *filter.AdminFilter) error {
	return a.dbDriver.Delete(ctx, &model.Admin{}, f)
}

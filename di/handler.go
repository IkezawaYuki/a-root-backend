package di

import (
	"IkezawaYuki/a-root-backend/interface/handler"
)

func NewCustomerHandler() handler.CustomerHandler {
	return handler.NewCustomerHandler(
		NewCustomerUsecase(),
	)
}

func NewAdminHandler() handler.AdminHandler {
	return handler.NewAdminHandler(
		NewAdminUsecase(),
	)
}

package di

import (
	"IkezawaYuki/a-root-backend/interface/handler"
)

func NewCustomerHandler() handler.CustomerHandler {
	return handler.NewCustomerHandler(
		NewCustomerUsecase(),
		redisClient,
	)
}

func NewAdminHandler() handler.AdminHandler {
	return handler.NewAdminHandler(
		NewAdminUsecase(),
		redisClient,
	)
}

func NewBatchHandler() handler.BatchHandler {
	return handler.NewBatchHandler(
		NewBatchUsecase(),
	)
}

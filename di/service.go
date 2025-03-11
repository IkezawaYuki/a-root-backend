package di

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/service"
)

func NewAdminService() service.AdminService {
	return service.NewAdminService(
		NewAdminRepository(),
	)
}

func NewPostService() service.PostService {
	return service.NewPostService(
		NewPostRepository(),
	)
}

func NewCustomerService() service.CustomerService {
	return service.NewCustomerService(
		NewCustomerRepository(),
	)
}

func NewAuthService() service.AuthService {
	return service.NewAuthService(
		NewRedisRepository(),
	)
}

func NewGraphAPI() service.GraphAPI {
	return service.NewGraph(
		infrastructure.NewHttpClient(),
	)
}

func NewFileTransfer() service.FileService {
	return service.NewFileService(
		infrastructure.NewHttpClient(),
	)
}

func NewSlackService() service.SlackService {
	return service.NewSlackService(
		infrastructure.NewHttpClient(),
	)
}

func NewOpenaiService() service.OpenaiService {
	return service.NewOpenaiService(
		infrastructure.NewOpenAI(config.Env.OpenAiApiKey),
		NewRedisRepository(),
	)
}

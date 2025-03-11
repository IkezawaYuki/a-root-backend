package di

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/interface/repository"
)

func NewBaseRepository() repository.BaseRepository {
	return repository.NewBaseRepository(
		NewDbDriver(),
	)
}

func NewAdminRepository() repository.AdminRepository {
	return repository.NewAdminRepository(
		NewDbDriver(),
	)
}

func NewPostRepository() repository.PostRepository {
	return repository.NewPostRepository(
		NewDbDriver(),
	)
}

func NewCustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(
		NewDbDriver(),
	)
}

func NewRodutRepository() repository.RodutRepository {
	return repository.NewRodutRepository(
		infrastructure.NewHttpClient(),
	)
}

func NewRedisRepository() repository.RedisRepository {
	return repository.NewRedisRepository(
		redisClient,
	)
}

func NewMailRepository() repository.MailRepository {
	return repository.NewMailRepository(
		NewMailDriver(),
		config.Env.FromEmail,
	)
}

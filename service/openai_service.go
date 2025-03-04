package service

import (
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/interface/dto/external"
	"IkezawaYuki/a-root-backend/interface/repository"
	"context"
)

type OpenaiService interface {
	Maika(ctx context.Context, dto external.OpenaiDto) (*external.OpenaiResult, error)
}

type openaiService struct {
	openai    infrastructure.OpenAI
	redisRepo repository.RedisRepository
}

func NewOpenaiService(
	openai infrastructure.OpenAI,
	redisRepo repository.RedisRepository,
) OpenaiService {
	return &openaiService{
		openai:    openai,
		redisRepo: redisRepo,
	}
}

func (o *openaiService) Maika(ctx context.Context, dto external.OpenaiDto) (*external.OpenaiResult, error) {

}

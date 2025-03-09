package usecase

import (
	"IkezawaYuki/a-root-backend/interface/dto/res"
	"IkezawaYuki/a-root-backend/service"
	"context"
	"sync"
)

type batchUsecase struct {
	customerService service.CustomerService
	customerUsecase CustomerUsecase
	slack           service.SlackService
}

type BatchUsecase interface {
	SyncInstagramToWordPress(ctx context.Context) (*res.Message, error)
	RefreshToken(ctx context.Context) (*res.Message, error)
}

func NewBatchUsecase(customerService service.CustomerService, customerUsecase CustomerUsecase, slackService service.SlackService) BatchUsecase {
	return &batchUsecase{
		customerService: customerService,
		customerUsecase: customerUsecase,
		slack:           slackService,
	}
}

const semaphore = 20

const (
	syncInstagramToWordPress = "インスタグラム => ワードプレス"
	refreshToken             = "トークン更新"
)

func (b *batchUsecase) SyncInstagramToWordPress(ctx context.Context) (*res.Message, error) {
	customers, err := b.customerService.FindAuthCustomers(ctx)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, semaphore) // 最大20件の並列処理を制限するためのセマフォ

	for _, customer := range customers {
		wg.Add(1)
		sem <- struct{}{}

		go func(customerID int) {
			defer wg.Done()
			defer func() { <-sem }()
			if _, err := b.customerUsecase.FetchAndPost(ctx, customerID); err != nil {
				_ = b.slack.Error(ctx, syncInstagramToWordPress, err)
				return
			}
		}(int(customer.ID))
	}

	wg.Wait()

	return &res.Message{Message: "ok"}, nil
}

func (b *batchUsecase) RefreshToken(ctx context.Context) (*res.Message, error) {
	customers, err := b.customerService.FindAuthCustomers(ctx)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, semaphore) // 最大20件の並列処理を制限するためのセマフォ

	for _, customer := range customers {
		wg.Add(1)
		sem <- struct{}{}

		go func(customerID int) {
			defer wg.Done()
			defer func() { <-sem }()
			if _, err := b.customerUsecase.RefreshToken(ctx, customerID); err != nil {
				_ = b.slack.Error(ctx, refreshToken, err)
				return
			}
		}(int(customer.ID))
	}

	wg.Wait()

	return &res.Message{Message: "ok"}, nil
}

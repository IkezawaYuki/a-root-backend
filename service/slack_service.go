package service

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/interface/dto/external"
	"context"
)

type SlackService interface {
	SendAlert(ctx context.Context, msg string) error
	SendNotification(ctx context.Context, msg string) error
}

type slackService struct {
	webhookURL string
	httpClient infrastructure.HttpClient
}

func NewSlackService(httpClient infrastructure.HttpClient) SlackService {
	return &slackService{
		webhookURL: config.Env.SlackWebhookURL,
		httpClient: httpClient,
	}
}

func (s *slackService) SendAlert(ctx context.Context, msg string) error {
	payload := external.SlackPayload{
		IconEmoji: ":wink",
		Text:      msg,
		Username:  "aroot",
	}
	_, err := s.httpClient.PostRequest(ctx, s.webhookURL, payload, "")
	if err != nil {
		return err
	}
	return nil
}

func (s *slackService) SendNotification(ctx context.Context, msg string) error {
	payload := external.SlackPayload{
		IconEmoji: ":wink",
		Text:      msg,
		Username:  "aroot",
	}
	_, err := s.httpClient.PostRequest(ctx, s.webhookURL, payload, "")
	if err != nil {
		return err
	}
	return nil
}

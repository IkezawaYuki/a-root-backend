package service

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/domain/model"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/interface/dto/external"
	"IkezawaYuki/a-root-backend/util"
	"context"
	"fmt"
)

type SlackService interface {
	Error(ctx context.Context, msg string, err error) error
	Log(ctx context.Context, customer *model.Customer, post *model.Post) error
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

func (s *slackService) send(ctx context.Context, msg string) error {
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

func (s *slackService) Error(ctx context.Context, method string, err error) error {
	trace := util.GetStackTrace()
	msg := fmt.Sprintf("```● %s\n%s\n%s```", method, err.Error(), trace)
	return s.send(ctx, msg)
}

func (s *slackService) Log(ctx context.Context, customer *model.Customer, post *model.Post) error {
	msg := fmt.Sprintf("```● %s\n%s\n%s```", customer.Name, post.InstagramMediaID, post.WordpressLink)
	return s.send(ctx, msg)
}

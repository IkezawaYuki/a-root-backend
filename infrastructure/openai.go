package infrastructure

import (
	"IkezawaYuki/a-root-backend/interface/dto/external"
	"context"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAI interface {
	Chat(ctx context.Context, dto external.OpenaiDto) (*external.OpenaiResult, error)
}

type openaiImpl struct {
	client *openai.Client
}

func NewOpenAI(apiKey string) OpenAI {
	return &openaiImpl{
		client: openai.NewClient(
			option.WithAPIKey(apiKey),
		),
	}
}

func (o *openaiImpl) Chat(ctx context.Context, dto external.OpenaiDto) (*external.OpenaiResult, error) {
	chatCompletion, err := o.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(dto.System),
			openai.UserMessage(dto.User),
		}),
		Model: openai.F(openai.ChatModelGPT4oMini),
	})
	if err != nil {
		return nil, err
	}
	return &external.OpenaiResult{
		Content: chatCompletion.Choices[0].Message.Content,
	}, nil
}

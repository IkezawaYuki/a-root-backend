package infrastructure

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailDriver interface {
	Send(from, to, subject, body string) error
	SendBulk(from string, to []string, subject, body string) error
}

func NewSendgridDriver(
	apiKey string,
	environment string,
) MailDriver {
	client := sendgrid.NewSendClient(apiKey)
	return &sendgridDriver{
		client:      client,
		environment: environment,
	}
}

type sendgridDriver struct {
	client      *sendgrid.Client
	environment string
}

func (d *sendgridDriver) Send(from, to, subject, body string) error {
	fromInfo := mail.NewEmail(from, from)
	toInfo := mail.NewEmail(to, to)

	message := mail.NewSingleEmail(fromInfo, subject, toInfo, body, "")

	message.AddCategories(fmt.Sprintf("service:a-root-%s", d.environment))
	message.SetCustomArg("env", d.environment)
	message.SetCustomArg("service", "a-root")

	_, err := d.client.Send(message)
	if err != nil {
		fmt.Println("failed to send mail")
		return err
	}
	return nil
}

func (d *sendgridDriver) SendBulk(from string, to []string, subject, body string) error {
	fromInfo := mail.NewEmail(from, from)
	personalizations := make([]*mail.Personalization, len(to))

	// 宛先リストに対してPersonalizationを作成
	for i, recipient := range to {
		toInfo := mail.NewEmail(recipient, recipient)
		personalization := mail.NewPersonalization()
		personalization.AddTos(toInfo)

		personalization.SetCustomArg("env", d.environment)
		personalization.SetCustomArg("service", "a-root")

		personalizations[i] = personalization
	}

	// メールメッセージを作成
	message := mail.NewSingleEmail(fromInfo, subject, nil, body, "")
	message.Personalizations = personalizations

	// 環境識別用のカテゴリを追加
	message.AddCategories(fmt.Sprintf("service:a-root-%s", d.environment))

	// メールを送信
	_, err := d.client.Send(message)
	if err != nil {
		fmt.Println("failed to send bulk mail:", err)
		return err
	}
	return nil
}

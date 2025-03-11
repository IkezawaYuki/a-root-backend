package repository

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/view"
	"bytes"
	"fmt"
	"html/template"
)

type MailRepository interface {
	TempRegisterMail(to string, token string) error
}

func NewMailRepository(driver infrastructure.MailDriver, from string) MailRepository {
	return &mailRepository{
		driver: driver,
		from:   from,
	}
}

type mailRepository struct {
	driver infrastructure.MailDriver
	from   string
}

const (
	TemplateRegisterMail = "view/mail/temp_register.txt"
)

func (r *mailRepository) TempRegisterMail(to string, token string) error {
	body, err := parseTemplate(TemplateRegisterMail, struct{ URL string }{
		URL: fmt.Sprintf("%s/confirm_mail?token=%s", config.Env.FrontendUrl, token),
	})
	if err != nil {
		return err
	}
	return r.driver.Send(r.from, to, "【A-Root】メールアドレス認証のお願い", body)
}

func parseTemplate(templatePath string, data interface{}) (string, error) {
	templateBytes, err := view.Root.ReadFile(templatePath)
	if err != nil {
		return "", err
	}
	t, err := template.New(templatePath).Parse(string(templateBytes))
	if err != nil {
		return "", err
	}
	buffer := &bytes.Buffer{}
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

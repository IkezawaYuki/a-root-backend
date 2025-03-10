package infrastructure

import (
	"fmt"
	"net/smtp"
	"strings"
)

func NewLocalMailDriver(
	hostname string,
	port int,
) MailDriver {
	return &localMailDriver{
		hostname: hostname,
		port:     port,
	}
}

type localMailDriver struct {
	hostname string
	port     int
}

func (d *localMailDriver) Send(from, to, subject, body string) error {
	msg := []byte(strings.ReplaceAll(
		fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join([]string{to}, ","), subject, body), "\n", "\r\n"),
	)
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", d.hostname, d.port), nil, from, []string{to}, msg); err != nil {
		return err
	}
	return nil
}

func (d *localMailDriver) SendBulk(from string, to []string, subject string, body string) error {
	msg := []byte(strings.ReplaceAll(
		fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(to, ","), subject, body), "\n", "\r\n"),
	)
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", d.hostname, d.port), nil, from, to, msg); err != nil {
		return err
	}
	return nil
}

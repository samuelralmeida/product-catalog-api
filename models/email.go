package models

import (
	"fmt"

	"github.com/go-mail/mail/v2"
	"github.com/samuelralmeida/product-catalog-api/env"
)

const (
	DefaultSender = "support@catalogapi.com"
)

type Email struct {
	From      string
	To        []string
	Subject   string
	Plaintext string
	HTML      string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailService(config *env.Config) *EmailService {
	es := &EmailService{
		dialer: mail.NewDialer(config.Smtp.Host, config.Smtp.Port, config.Smtp.Username, config.Smtp.Password),
	}
	return es
}

type EmailService struct {
	DefaultSender string

	dialer *mail.Dialer
}

func (es *EmailService) Send(email Email) error {
	msg := mail.NewMessage()

	msg.SetHeader("To", email.To...)
	msg.SetHeader("Subject", email.Subject)

	es.setFrom(msg, email)

	switch {
	case email.Plaintext != "" && email.HTML != "":
		msg.SetBody("text/plain", email.Plaintext)
		msg.AddAlternative("text/html", email.HTML)
	case email.Plaintext != "":
		msg.SetBody("text/plain", email.Plaintext)
	case email.HTML != "":
		msg.SetBody("text/html", email.HTML)
	}

	err := es.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send email: %w", err)
	}
	return nil
}

func (es *EmailService) setFrom(msg *mail.Message, email Email) {
	var from string
	switch {
	case email.From != "":
		from = email.From
	case es.DefaultSender != "":
		from = es.DefaultSender
	default:
		from = DefaultSender
	}
	msg.SetHeader("From", from)
}

func (es *EmailService) ForgotPassword(to, resetUrl string) error {
	email := Email{
		Subject:   "Reset your password",
		To:        []string{to},
		Plaintext: "To reset your password, please visit the following link: " + resetUrl,
		HTML:      `<p>To reset your password, please visit the following link: <a href="` + resetUrl + `">` + resetUrl + `</a></p>`,
	}
	err := es.Send(email)
	if err != nil {
		return fmt.Errorf("forgot password email: %w", err)
	}
	return nil
}

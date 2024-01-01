package email

import (
	"context"
	"fmt"
	"io"

	"github.com/go-mail/mail/v2"
)

const (
	defaultSender = "support@catalogapi.com"
)

type email struct {
	From      string
	To        []string
	Subject   string
	Plaintext string
	HTML      string
}

type UseCases struct {
	Dialer *mail.Dialer
	Writer io.Writer
}

func (uc *UseCases) SendForgotPassword(ctx context.Context, to string, resetUrl string) error {
	mail := email{
		Subject:   "Reset your password",
		To:        []string{to},
		Plaintext: "To reset your password, please visit the following link: " + resetUrl,
		HTML:      `<p>To reset your password, please visit the following link: <a href="` + resetUrl + `">` + resetUrl + `</a></p>`,
	}

	err := uc.send(mail)
	if err != nil {
		return fmt.Errorf("forgot password email: %w", err)
	}
	return nil
}

func (uc *UseCases) send(email email) error {
	msg := mail.NewMessage()

	msg.SetHeader("To", email.To...)
	msg.SetHeader("Subject", email.Subject)

	setFrom(msg, email)

	switch {
	case email.Plaintext != "" && email.HTML != "":
		msg.SetBody("text/plain", email.Plaintext)
		msg.AddAlternative("text/html", email.HTML)
	case email.Plaintext != "":
		msg.SetBody("text/plain", email.Plaintext)
	case email.HTML != "":
		msg.SetBody("text/html", email.HTML)
	}

	if uc.Writer != nil {
		msg.WriteTo(uc.Writer)
	}

	err := uc.Dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send email: %w", err)
	}
	return nil
}

func setFrom(msg *mail.Message, email email) {
	var from string
	switch {
	case email.From != "":
		from = email.From
	default:
		from = defaultSender
	}
	msg.SetHeader("From", from)
}

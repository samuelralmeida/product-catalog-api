package email

import (
	"fmt"

	"github.com/go-mail/mail/v2"
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

type Dialer struct {
	dialer *mail.Dialer
}

func NewDial(dialer *mail.Dialer) *Dialer {
	return &Dialer{dialer: dialer}
}

func (d *Dialer) ForgotPassword(to, resetUrl string) error {
	email := Email{
		Subject:   "Reset your password",
		To:        []string{to},
		Plaintext: "To reset your password, please visit the following link: " + resetUrl,
		HTML:      `<p>To reset your password, please visit the following link: <a href="` + resetUrl + `">` + resetUrl + `</a></p>`,
	}
	err := d.Send(email)
	if err != nil {
		return fmt.Errorf("forgot password email: %w", err)
	}
	return nil
}

func (d *Dialer) Send(email Email) error {
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

	err := d.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send email: %w", err)
	}
	return nil
}

func setFrom(msg *mail.Message, email Email) {
	var from string
	switch {
	case email.From != "":
		from = email.From
	default:
		from = DefaultSender
	}
	msg.SetHeader("From", from)
}

package mailtrap

import (
	"github.com/go-mail/mail/v2"
	"github.com/samuelralmeida/product-catalog-api/internal/env"
)

func NewDialer(config *env.Config) *mail.Dialer {
	return mail.NewDialer(
		config.Smtp.Host,
		config.Smtp.Port,
		config.Smtp.Username,
		config.Smtp.Password,
	)
}

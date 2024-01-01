package service

import (
	"context"
	"fmt"
	"net/url"

	"github.com/samuelralmeida/product-catalog-api/entity"
	"github.com/samuelralmeida/product-catalog-api/env"
)

type UserUseCases interface {
	Create(ctx context.Context, email, password string) (*entity.User, error)
	Authenticate(ctx context.Context, email, password string) (*entity.User, error)
	CreatePasswordResetToken(ctx context.Context, email string) (*entity.PasswordReset, error)
	UserByPasswordResetToken(ctx context.Context, passwordResetToken string) (*entity.User, error)
	UpdatePassword(ctx context.Context, userID uint, password string) error
}

type SessionUseCases interface {
	Create(ctx context.Context, user *entity.User) (*entity.Session, error)
	User(ctx context.Context, sessionToken string) (*entity.User, error)
	Delete(ctx context.Context, sessionToken string) error
}

type MailUseCases interface {
	SendForgotPassword(ctx context.Context, to string, resetUrl string) error
}

type UserService struct {
	UserUseCases    UserUseCases
	SessionUseCases SessionUseCases
	MailUseCase     MailUseCases
	Config          *env.Config
}

func (us *UserService) Create(ctx context.Context, email, password string) (*entity.User, *entity.Session, error) {
	user, err := us.UserUseCases.Create(ctx, email, password)
	if err != nil {
		return nil, nil, fmt.Errorf("create user: %w", err)
	}
	session, err := us.SessionUseCases.Create(ctx, user)
	if err != nil {
		return user, nil, fmt.Errorf("create session")
	}
	return user, session, nil
}

func (us *UserService) Autheticate(ctx context.Context, email string, password string) (*entity.Session, error) {
	user, err := us.UserUseCases.Authenticate(ctx, email, password)
	if err != nil {
		return nil, fmt.Errorf("authenticate user: %w", err)
	}

	session, err := us.SessionUseCases.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	return session, nil
}

func (us *UserService) SignOut(ctx context.Context, sessionToken string) error {
	err := us.SessionUseCases.Delete(ctx, sessionToken)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	return nil
}

func (us *UserService) ForgotPassword(ctx context.Context, email string) (*entity.PasswordReset, error) {
	pwReset, err := us.UserUseCases.CreatePasswordResetToken(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("create password reset token")
	}

	vals := url.Values{"token": {pwReset.Token}}
	resetUrl := fmt.Sprintf("%s/reset-pw?%s", us.Config.Server.Url, vals.Encode())

	err = us.MailUseCase.SendForgotPassword(ctx, email, resetUrl)
	if err != nil {
		return nil, fmt.Errorf("send forgot password")
	}

	return pwReset, nil
}

func (us *UserService) ResetPassword(ctx context.Context, resetPasswordToken string, password string) (*entity.Session, error) {
	user, err := us.UserUseCases.UserByPasswordResetToken(ctx, resetPasswordToken)
	if err != nil {
		return nil, fmt.Errorf("reset password: %w", err)
	}

	err = us.UserUseCases.UpdatePassword(ctx, user.ID, password)
	if err != nil {
		return nil, fmt.Errorf("update password: %w", err)
	}

	session, err := us.SessionUseCases.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	return session, nil
}

func (us *UserService) User(ctx context.Context, sessionToken string) (*entity.User, error) {
	user, err := us.SessionUseCases.User(ctx, sessionToken)
	if err != nil {
		return nil, fmt.Errorf("user by session token: %w", err)
	}
	return user, nil
}

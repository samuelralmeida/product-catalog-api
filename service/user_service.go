package service

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/entity"
)

type UserUseCases interface {
	Create(ctx context.Context, email, password string) (*entity.User, error)
	Authenticate(ctx context.Context, email, password string) (*entity.User, error)
}

type SessionUseCases interface {
	Create(ctx context.Context, user *entity.User) (*entity.Session, error)
	User(ctx context.Context, sessionToken string) (*entity.User, error)
	Delete(ctx context.Context, sessionToken string) error
}

type UserService struct {
	userUseCases    UserUseCases
	sessionUseCases SessionUseCases
}

func (us *UserService) Create(ctx context.Context, email, password string) (*entity.Session, error) {
	user, err := us.userUseCases.Create(ctx, email, password)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	session, err := us.sessionUseCases.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create session")
	}
	return session, nil
}

func (us *UserService) Signin(ctx context.Context, email, password string) (*entity.Session, error) {
	user, err := us.userUseCases.Authenticate(ctx, email, password)
	if err != nil {
		return nil, fmt.Errorf("siging: %w", err)
	}
	session, err := us.sessionUseCases.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create session")
	}
	return session, nil
}

func (us *UserService) UserBySessionToken(ctx context.Context, sessionToken string) (*entity.User, error) {
	user, err := us.sessionUseCases.User(ctx, sessionToken)
	if err != nil {
		return nil, fmt.Errorf("get user by token session: %w", err)
	}
	return user, nil
}

func (us *UserService) Signout(ctx context.Context, sessionToken string) error {
	err := us.sessionUseCases.Delete(ctx, sessionToken)
	if err != nil {
		return fmt.Errorf("signout: %w", err)
	}
	return nil
}

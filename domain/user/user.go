package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/samuelralmeida/product-catalog-api/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Save(ctx context.Context, email, passwordHash string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type UserUseCases struct {
	Repository UserRepository
}

func (uuc *UserUseCases) Create(ctx context.Context, email, password string) (*entity.User, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generate password hash: %w", err)
	}

	passwordHash := string(hashedBytes)
	user, err := uuc.Repository.Save(ctx, strings.ToLower(email), passwordHash)
	if err != nil {
		return nil, fmt.Errorf("save user: %w", err)
	}
	return user, nil
}

func (uuc *UserUseCases) Authenticate(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := uuc.Repository.GetByEmail(ctx, strings.ToLower(email))
	if err != nil {
		return nil, fmt.Errorf("autheticate user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password invalid: %w", err)
	}

	return user, nil
}

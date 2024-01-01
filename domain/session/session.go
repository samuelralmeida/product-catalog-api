package session

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/entity"
	"github.com/samuelralmeida/product-catalog-api/internal/hash"
	"github.com/samuelralmeida/product-catalog-api/internal/rand"
)

const (
	MinBytesPerToken = 32
)

type SessionRepository interface {
	Save(ctx context.Context, user *entity.User, session *entity.Session) error
	GetUserByTokenHash(ctx context.Context, tokenHash string) (*entity.User, error)
	Delete(ctx context.Context, tokenHash string) error
}

type UseCases struct {
	Repository SessionRepository
	// BytesPerToken is used to determine ho wmany bytes to use when generating
	// each sessio token. If this value is not set or is less than the
	// MinBytesPerToken it will be ignored and MinBytesPerToken will be used
	BytesPerToken int
}

func (suc *UseCases) Create(ctx context.Context, user *entity.User) (*entity.Session, error) {
	bytesPerToken := suc.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}

	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create session token: %w", err)
	}

	session := &entity.Session{
		UserID:    user.ID,
		Token:     token,
		TokenHash: hash.Sha256(token),
	}

	err = suc.Repository.Save(ctx, user, session)
	if err != nil {
		return nil, fmt.Errorf("save session token hash: %w", err)
	}
	return session, err
}

func (suc *UseCases) User(ctx context.Context, token string) (*entity.User, error) {
	tokenHash := hash.Sha256(token)
	return suc.Repository.GetUserByTokenHash(ctx, tokenHash)
}

func (suc *UseCases) Delete(ctx context.Context, token string) error {
	tokenHash := hash.Sha256(token)
	return suc.Repository.Delete(ctx, tokenHash)
}

package session

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/entity"
	"github.com/samuelralmeida/product-catalog-api/internal/rand"
)

const (
	MinBytesPerToken = 32
)

type SessionRepository interface {
	Save(ctx context.Context, user *entity.User, session *entity.Session) error
	GetUserByToken(ctx context.Context, tokenHash string) (*entity.User, error)
	Delete(ctx context.Context, tokenHash string) error
}

type SessionUserCases struct {
	Repository SessionRepository
	// BytesPerToken is used to determine ho wmany bytes to use when generating
	// each sessio token. If this value is not set or is less than the
	// MinBytesPerToken it will be ignored and MinBytesPerToken will be used
	BytesPerToken int
}

func (suc *SessionUserCases) Create(ctx context.Context, user *entity.User) (*entity.Session, error) {
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
		TokenHash: hash(token),
	}

	err = suc.Repository.Save(ctx, user, session)
	if err != nil {
		return nil, fmt.Errorf("save session token hash: %w", err)
	}
	return session, err
}

func (suc *SessionUserCases) User(ctx context.Context, token string) (*entity.User, error) {
	tokenHash := hash(token)
	return suc.Repository.GetUserByToken(ctx, tokenHash)
}

func (suc *SessionUserCases) Delete(ctx context.Context, token string) error {
	tokenHash := hash(token)
	return suc.Repository.Delete(ctx, tokenHash)
}

func hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

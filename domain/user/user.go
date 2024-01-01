package user

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/samuelralmeida/product-catalog-api/entity"
	"github.com/samuelralmeida/product-catalog-api/internal/hash"
	"github.com/samuelralmeida/product-catalog-api/internal/rand"
	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultResetPasswordTokendDuration = 1 * time.Hour
	MinBytesPerResetPasswordToken      = 32
)

type UserRepository interface {
	Save(ctx context.Context, email, passwordHash string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	CreatePasswordResetHash(ctx context.Context, pwReset *entity.PasswordReset) error
	ConsumePasswordResetHash(ctx context.Context, passwordResetHash string) (*entity.User, *entity.PasswordReset, error)
	DeletePasswordResetHash(ctx context.Context, passwordResetID uint) error
	UpdatePassword(ctx context.Context, userID uint, passwordHash string) error
}

type UseCases struct {
	Repository                 UserRepository
	BytesPerResetPasswordToken int
	ResetPasswordTokenDuration time.Duration
}

func (uuc *UseCases) Create(ctx context.Context, email, password string) (*entity.User, error) {
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

func (uuc *UseCases) Authenticate(ctx context.Context, email, password string) (*entity.User, error) {
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

func (uuc *UseCases) CreatePasswordResetToken(ctx context.Context, email string) (*entity.PasswordReset, error) {
	email = strings.ToLower(email)
	user, err := uuc.Repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("get by email to reset passord token: %w", err)
	}

	bytesPerToken := uuc.BytesPerResetPasswordToken
	if bytesPerToken < MinBytesPerResetPasswordToken {
		bytesPerToken = MinBytesPerResetPasswordToken
	}

	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	duration := uuc.ResetPasswordTokenDuration
	if duration == 0 {
		duration = DefaultResetPasswordTokendDuration
	}

	pwReset := &entity.PasswordReset{
		UserID:    user.ID,
		Token:     token,
		TokenHash: hash.Sha256(token),
		ExpiresAt: time.Now().Add(duration),
	}

	err = uuc.Repository.CreatePasswordResetHash(ctx, pwReset)
	if err != nil {
		return nil, fmt.Errorf("create password reset token: %w", err)
	}
	return pwReset, nil
}

func (uuc *UseCases) UserByPasswordResetToken(ctx context.Context, passwordResetToken string) (*entity.User, error) {
	tokenHash := hash.Sha256(passwordResetToken)

	user, pwReset, err := uuc.Repository.ConsumePasswordResetHash(ctx, tokenHash)
	if err != nil {
		return nil, fmt.Errorf("consume password reset token: %w", err)
	}

	if time.Now().After(pwReset.ExpiresAt) {
		return nil, fmt.Errorf("token expires - %s", passwordResetToken)
	}

	err = uuc.Repository.DeletePasswordResetHash(ctx, pwReset.ID)
	if err != nil {
		log.Println("delete password reset token")
	}

	return user, nil
}

func (uuc *UseCases) UpdatePassword(ctx context.Context, userID uint, password string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("generate password hash: %w", err)
	}

	passwordHash := string(hashedBytes)
	err = uuc.Repository.UpdatePassword(ctx, userID, passwordHash)
	if err != nil {
		return fmt.Errorf("update password: %w", err)
	}
	return nil
}

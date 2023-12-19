package models

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/internal/rand"
)

const (
	MinBytesPerToken = 32
)

// TODO: separar struct de session da aplicação da session do repositório
// não faz sentido a session da aplicação ter um id do banco,
// ela pode vir de outros lugares
type Session struct {
	ID     uint
	UserID uint
	// Token is only set when creating a new session. When looking up a session
	// this will be left empty, as we only store the hash of a session token
	// in our database and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB database.Database
	// BytesPerToken is used to determine ho wmany bytes to use when generating
	// each sessio token. If this value is not set or is less than the
	// MinBytesPerToken it will be ignored and MinBytesPerToken will be used
	BytesPerToken int
}

// Create will create a new session for the user provided. The session token
// will be returned as the Token field on the Session type, but only the hashed
// session token is stored in the database.
func (ss *SessionService) Create(userID uint) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}

	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create sessoion token: %w", err)
	}

	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}

	row := ss.DB.QueryRowContext(context.Background(), insertOrUpdateSessionQuery, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err != nil {
		return nil, fmt.Errorf("update session hash: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)

	var user User
	row := ss.DB.QueryRowContext(context.Background(), selectSessionQuery, tokenHash)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("select session: %w", err)
	}

	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.ExecContext(context.Background(), deleteSessionQuery, tokenHash)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	return nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

const (
	insertOrUpdateSessionQuery = "INSERT INTO sessions (user_id, token_hash) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET token_hash = $2 RETURNING id"
	selectSessionQuery         = "SELECT u.id, u.email, u.password_hash FROM users u JOIN sessions s ON u.id = s.user_id WHERE s.token_hash = $1"
	deleteSessionQuery         = "DELETE FROM sessions WHERE token_hash = $1"
)

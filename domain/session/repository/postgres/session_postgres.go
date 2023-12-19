package sessionpostgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/entity"
)

type SessionRepository struct {
	DB database.Database
}

const (
	updateSessionQuery      = "UPDATE sessions SET token_hash = $2 WHERE user_id = $1 RETURNING id"
	insertSessionQuery      = "INSERT INTO sessions (user_id, token_hash) VALUES ($1, $2) RETURNING id"
	deleteSessionQuery      = "DELETE FROM sessions WHERE token_hash = $1"
	getUserByTokenHashQUery = "SELECT id, email, password_hash FROM users u JOIN sessions s ON s.user_id = u.id WHERE s.token_hash = $1"
)

func (sr *SessionRepository) Save(ctx context.Context, user *entity.User, session *entity.Session) error {
	row := sr.DB.QueryRowContext(ctx, updateSessionQuery, user.ID, session.TokenHash)
	var id int
	err := row.Scan(id)
	if err == sql.ErrNoRows {
		row = sr.DB.QueryRowContext(context.Background(), insertSessionQuery, session.UserID, session.TokenHash)
		err = row.Scan(id)
		if err != nil {
			return fmt.Errorf("insert session hash: %w", err)
		}
	}
	if err != nil {
		return fmt.Errorf("update session hash: %w", err)
	}
	return nil
}

func (sr *SessionRepository) GetUserByTokenHash(ctx context.Context, tokenHash string) (*entity.User, error) {
	var user entity.User
	row := sr.DB.QueryRowContext(ctx, getUserByTokenHashQUery, tokenHash)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get user by token hash: %w", err)
	}
	return &user, nil
}

func (sr *SessionRepository) Delete(ctx context.Context, tokenHash string) error {
	_, err := sr.DB.ExecContext(ctx, deleteSessionQuery, tokenHash)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	return nil
}

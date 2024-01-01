package userpostgres

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/entity"
)

type UserRepository struct {
	DB database.Database
}

const insertUserQuery = `INSERT INTO users (email, password_hash) values ($1, $2) RETURNING id`

func (ur *UserRepository) Save(ctx context.Context, email, passwordHash string) (*entity.User, error) {
	user := entity.User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	row := ur.DB.QueryRowContext(ctx, insertUserQuery, email, passwordHash)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &user, nil
}

const selectUserByEmailQuery = `SELECT id, password_hash FROM users WHERE email = $1`

// TODO: get pelo hash do password pode ser mais interessante e seguro
func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := entity.User{
		Email: email,
	}
	row := ur.DB.QueryRowContext(ctx, selectUserByEmailQuery, user.Email)
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return &user, nil
}

const insertPasswordResetHashQuery = `
	INSERT INTO password_resets (user_id, token_hash, expires_at)
	VALUES ($1, $2, $3) ON CONFLICT (user_id) DO
	UPDATE
	SET token_hash = $2, expires_at = $3
	RETURNING id;
`

func (ur *UserRepository) CreatePasswordResetHash(ctx context.Context, pwReset *entity.PasswordReset) error {
	row := ur.DB.QueryRowContext(
		ctx,
		insertPasswordResetHashQuery,
		pwReset.UserID, pwReset.TokenHash, pwReset.ExpiresAt,
	)
	err := row.Scan(&pwReset.ID)
	if err != nil {
		return fmt.Errorf("insert password reset hash: %w", err)
	}
	return nil
}

const selectByPaswordResetHashQuery = `
	SELECT
		password_resets.id, password_resets.expires_at,
		users.id, users.email, users.password_hash
	FROM password_resets
	JOIN users ON users.id = password_resets.user_id
	WHERE password_resets.token_hash = $1;
`

func (ur *UserRepository) ConsumePasswordResetHash(ctx context.Context, passwordResetHash string) (*entity.User, *entity.PasswordReset, error) {
	var user entity.User
	var pwReset entity.PasswordReset
	row := ur.DB.QueryRowContext(ctx, selectByPaswordResetHashQuery, passwordResetHash)
	err := row.Scan(
		&pwReset.ID, &pwReset.ExpiresAt,
		&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, nil, fmt.Errorf("consume password reset hash: %w", err)
	}
	return &user, &pwReset, nil
}

const deletePasswordHashQuery = `DELETE FROM password_resets WHERE id = $1;`

func (ur *UserRepository) DeletePasswordResetHash(ctx context.Context, passwordResetID uint) error {
	_, err := ur.DB.ExecContext(ctx, deletePasswordHashQuery, passwordResetID)
	if err != nil {
		return fmt.Errorf("delete password reset hash: %w", err)
	}
	return nil
}

const updateUserPasswordQuery = `UPDATE users SET password_hash = $1 WHERE id = $2;`

func (ur *UserRepository) UpdatePassword(ctx context.Context, userID uint, passwordHash string) error {
	_, err := ur.DB.ExecContext(ctx, updateUserPasswordQuery, passwordHash, userID)
	if err != nil {
		return fmt.Errorf("update password: %w", err)
	}
	return nil
}

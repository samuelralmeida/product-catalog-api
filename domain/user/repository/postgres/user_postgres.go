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

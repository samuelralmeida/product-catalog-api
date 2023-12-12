package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/samuelralmeida/product-catalog-api/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint
	Email        string
	PasswordHash string
}

type UserService struct {
	DB database.Database
}

const insertUserQuery = `INSERT INTO users (email, password_hash) values ($1, $2) RETURNING id`

func (us *UserService) Create(email, password string) (*User, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generate password hash: %w", err)
	}

	passwordHash := string(hashedBytes)

	user := User{
		Email:        strings.ToLower(email),
		PasswordHash: passwordHash,
	}

	tx, err := us.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("beging transaction: %w", err)
	}

	row := tx.QueryRowContext(context.Background(), insertUserQuery, email, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("create user: %w", err)
	}
	tx.Commit()

	return &user, nil
}

func (us *UserService) Update(user *User) error {
	return nil
}

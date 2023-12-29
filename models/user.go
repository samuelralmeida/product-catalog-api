package models

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/samuelralmeida/product-catalog-api/database"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailTaken = errors.New("models: email address is already in use")
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

	row := us.DB.QueryRowContext(context.Background(), insertUserQuery, email, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			println(pgError.Code, pgerrcode.UniqueViolation)
			if pgError.Code == pgerrcode.UniqueViolation {
				return nil, ErrEmailTaken
			}
		}
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

const selectUserByEmailQuery = `SELECT id, password_hash FROM users WHERE email = $1`

func (us *UserService) Authenticate(email, password string) (*User, error) {
	user := User{
		Email: strings.ToLower(email),
	}

	row := us.DB.QueryRowContext(context.Background(), selectUserByEmailQuery, user.Email)
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authenticate user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password invalid: %w", err)
	}

	return &user, nil
}

func (us *UserService) UpdatePassword(userID uint, password string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("update password: %w", err)
	}
	passwordHash := string(hashedBytes)
	_, err = us.DB.ExecContext(context.Background(), `
		UPDATE users
		SET password_hash = $2
		WHERE id = $1;
	`, userID, passwordHash)
	if err != nil {
		return fmt.Errorf("update password: %w", err)
	}
	return nil
}

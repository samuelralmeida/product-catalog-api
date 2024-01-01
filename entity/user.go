package entity

import "time"

type User struct {
	ID           uint
	Email        string
	PasswordHash string
}

type Session struct {
	UserID    uint
	Token     string
	TokenHash string
}

type PasswordReset struct {
	ID     uint
	UserID uint
	// Token is only set when a PasswordReset is being created.
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

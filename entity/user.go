package entity

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

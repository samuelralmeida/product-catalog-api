package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

type smtp struct {
	Host     string
	Port     int
	Username string
	Password string
}

type csrf struct {
	Key    string
	Secure bool
}

type server struct {
	Port string
	Url  string
}

type Config struct {
	Postgres postgres
	Smtp     smtp
	Csrf     csrf
	Server   server
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("load .env file: %w", err))
	}

	config := &Config{
		Postgres: postgres{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DATABASE"),
			SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		},
		Smtp: smtp{
			Host:     os.Getenv("SMTP_HOST"),
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Port:     parseIntEnv(os.Getenv("SMTP_PORT")),
		},
		Csrf: csrf{
			Key:    os.Getenv("CSRF_KEY"),
			Secure: false,
		},
		Server: server{
			Port: os.Getenv("SERVER_PORT"),
			Url:  os.Getenv("SERVER_URL"),
		},
	}

	return config
}

func parseIntEnv(env string) int {
	value, err := strconv.Atoi(env)
	if err != nil {
		panic(fmt.Errorf("parse env int (%s): %w", env, err))
	}
	return value
}

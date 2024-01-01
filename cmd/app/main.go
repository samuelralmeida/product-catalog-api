package main

import (
	"log"
	"os"

	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/database/postgres"
	"github.com/samuelralmeida/product-catalog-api/domain/session"
	"github.com/samuelralmeida/product-catalog-api/domain/session/repository/sessionpostgres"
	"github.com/samuelralmeida/product-catalog-api/domain/user"
	"github.com/samuelralmeida/product-catalog-api/domain/user/repository/userpostgres"
	"github.com/samuelralmeida/product-catalog-api/email"
	"github.com/samuelralmeida/product-catalog-api/email/mailtrap"
	"github.com/samuelralmeida/product-catalog-api/env"
	"github.com/samuelralmeida/product-catalog-api/internal/api"
	"github.com/samuelralmeida/product-catalog-api/internal/controllers"
	"github.com/samuelralmeida/product-catalog-api/internal/controllers/chi"
	"github.com/samuelralmeida/product-catalog-api/service"
)

func main() {
	config := env.Load()

	// database

	conn, err := postgres.Open(postgres.EnvConfig(config))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	db := database.NewDB(conn)

	// external services

	dialer := mailtrap.NewDialer(config)
	mailUseCases := email.UseCases{
		Dialer: dialer,
		Writer: os.Stdout, // comment to actually send email
	}

	// repositories

	userRepository := &userpostgres.UserRepository{DB: db}
	sessionRepository := &sessionpostgres.SessionRepository{DB: db}

	// use cases

	userUseCase := &user.UseCases{Repository: userRepository}
	sessionUseCase := &session.UseCases{Repository: sessionRepository}

	// services

	userService := &service.UserService{
		UserUseCases:    userUseCase,
		SessionUseCases: sessionUseCase,
		MailUseCase:     &mailUseCases,
		Config:          config,
	}

	// controller

	controller := &controllers.Controller{Config: config, UserService: userService}

	// template

	htmlTemplates := controllers.MustParseTemplates()

	// handler

	r := chi.Handlers(controller, htmlTemplates)

	err = api.Start(config, r)
	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}

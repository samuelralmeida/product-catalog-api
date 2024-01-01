package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/samuelralmeida/product-catalog-api/controllers"
	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/database/postgres"
	"github.com/samuelralmeida/product-catalog-api/env"
	"github.com/samuelralmeida/product-catalog-api/models"
	"github.com/samuelralmeida/product-catalog-api/templates"
	"github.com/samuelralmeida/product-catalog-api/views"
)

func main() {
	config := env.Load()

	conn, err := postgres.Open(postgres.EnvConfig(config))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	db := database.NewDB(conn)
	userService := models.UserService{DB: db}
	sessionService := models.SessionService{DB: db}
	emailService := models.NewEmailService(config)
	passwordResetService := models.PasswordResetService{DB: db}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	csrfMiddleware := csrf.Protect([]byte(config.Csrf.Key))
	r.Use(csrfMiddleware)

	umw := controllers.UserMiddleware{SessionService: &sessionService}
	r.Use(umw.SetUser)

	var tpl views.Template

	tpl = views.MustParseFS(templates.FS, "layout-page.gohtml", "home.gohtml")
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.MustParseFS(templates.FS, "layout-page.gohtml", "product.gohtml")
	r.Get("/product", controllers.ProductHandler(tpl))

	userController := controllers.Users{UserService: &userService, SessionService: &sessionService, EmailService: emailService, PasswordResetService: &passwordResetService}
	userController.Templates.Signup = views.MustParseFS(templates.FS, "layout-page.gohtml", "signup.gohtml")
	userController.Templates.Signin = views.MustParseFS(templates.FS, "layout-page.gohtml", "signin.gohtml")
	userController.Templates.ForgotPassword = views.MustParseFS(templates.FS, "layout-page.gohtml", "forgot-pw.gohtml")
	userController.Templates.CheckYourEmail = views.MustParseFS(templates.FS, "layout-page.gohtml", "check-your-email.gohtml")
	userController.Templates.ResetPassword = views.MustParseFS(templates.FS, "layout-page.gohtml", "reset-pw.gohtml")

	r.Get("/signup", userController.SignUp)
	r.Post("/users", userController.Create)
	r.Get("/signin", userController.SignIn)
	r.Post("/signin", userController.ProcessSignIn)
	r.Post("/signout", userController.ProcessSignOut)
	r.Get("/forgot-pw", userController.ForgotPassword)
	r.Post("/forgot-pw", userController.ProcessForgotPassword)
	r.Get("/reset-pw", userController.ResetPassword)
	r.Post("/reset-pw", userController.ProcessResetPassword)

	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", userController.CurrentUser)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })

	fmt.Printf("Starting the server on %s...\n", config.Server.Port)
	err = http.ListenAndServe(config.Server.Port, r)
	if err != nil {
		panic(err)
	}
}

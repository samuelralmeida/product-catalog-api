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
	"github.com/samuelralmeida/product-catalog-api/models"
	"github.com/samuelralmeida/product-catalog-api/templates"
	"github.com/samuelralmeida/product-catalog-api/views"
)

func main() {
	conn, err := postgres.Open(postgres.DefaultConfig())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	db := database.NewDB(conn)
	userService := models.UserService{DB: db}
	sessionService := models.SessionService{DB: db}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
	r.Use(csrfMiddleware)

	umw := controllers.UserMiddleware{SessionService: &sessionService}
	r.Use(umw.SetUser)

	var tpl views.Template

	tpl = views.MustParseFS(templates.FS, "layout-page.gohtml", "home.gohtml")
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.MustParseFS(templates.FS, "layout-page.gohtml", "product.gohtml")
	r.Get("/product", controllers.ProductHandler(tpl))

	userController := controllers.User{UserService: &userService, SessionService: &sessionService}
	userController.Templates.Signup = views.MustParseFS(templates.FS, "layout-page.gohtml", "signup.gohtml")
	userController.Templates.Signin = views.MustParseFS(templates.FS, "layout-page.gohtml", "signin.gohtml")

	r.Get("/signup", userController.SignUp)
	r.Post("/users", userController.Create)
	r.Get("/signin", userController.SignIn)
	r.Post("/signin", userController.ProcessSignIn)
	r.Post("/signout", userController.ProcessSignOut)

	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", userController.CurrentUser)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	us := models.UserService{DB: db}

	user, err := us.Create("samuel2@email.com", "password")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	var tpl views.Template

	tpl = views.MustParseFS(templates.FS, "layout-page.gohtml", "home.gohtml")
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.MustParseFS(templates.FS, "layout-page.gohtml", "product.gohtml")
	r.Get("/product", controllers.ProductHandler(tpl))

	userController := controllers.User{}
	userController.Templates.New = views.MustParseFS(templates.FS, "layout-page.gohtml", "signup.gohtml")

	r.Get("/signup", userController.New)
	r.Post("/users", userController.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

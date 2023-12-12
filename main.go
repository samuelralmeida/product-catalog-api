package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samuelralmeida/product-catalog-api/controllers"
	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/templates"
	"github.com/samuelralmeida/product-catalog-api/views"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := database.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		Database: "catalog",
		SSLMode:  "disable",
	}

	conn, err := sql.Open("pgx", cfg.PostgresUrl())
	if err != nil {
		panic(err)
	}

	db := database.NewDB(conn)

	err = db.PingContext(context.Background())
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

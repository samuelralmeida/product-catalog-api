package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samuelralmeida/product-catalog-api/controllers"
	"github.com/samuelralmeida/product-catalog-api/templates"
	"github.com/samuelralmeida/product-catalog-api/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	var tpl views.Template

	tpl = views.MustParseFS(templates.FS, "home.gohtml")
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.MustParseFS(templates.FS, "product.gohtml")
	r.Get("/product", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

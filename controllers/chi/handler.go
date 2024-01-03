package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/samuelralmeida/product-catalog-api/controllers"
)

func Handlers(controller *controllers.Controller, templates controllers.HtmlTemplates) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	csrfMiddleware := csrf.Protect([]byte(controller.Config.Csrf.Key))

	umw := controllers.UserMiddleware{UserService: controller.UserService}
	r.Use(umw.SetUser)

	userHandler := Users{
		UserService: controller.UserService,
		Templates:   templates,
	}

	productHandler := Products{
		ProductService: controller.ProductService,
	}

	r.Route("/users", func(r chi.Router) {
		r.Use(csrfMiddleware)

		r.Get("/", userHandler.Home)
		r.Post("/", userHandler.Create)
		r.Get("/signup", userHandler.SignUp)
		r.Get("/signin", userHandler.SignIn)
		r.Post("/signin", userHandler.ProcessSignIn)
		r.Post("/signout", userHandler.ProcessSignOut)
		r.Get("/forgot-pw", userHandler.ForgotPassword)
		r.Post("/forgot-pw", userHandler.ProcessForgotPassword)
		r.Get("/reset-pw", userHandler.ResetPassword)
		r.Post("/reset-pw", userHandler.ProcessResetPassword)

		r.Group(func(r chi.Router) {
			r.Use(umw.RequireUser)
			r.Get("/me", userHandler.CurrentUser)
		})

	})

	r.Route("/products", func(r chi.Router) {
		r.Get("/", productHandler.List)
		r.Post("/", productHandler.Create)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	return r
}

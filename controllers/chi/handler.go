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
	r.Use(csrfMiddleware)

	umw := controllers.UserMiddleware{UserService: controller.UserService}
	r.Use(umw.SetUser)

	userHandler := Users{
		UserService: controller.UserService,
		Templates:   templates,
	}

	r.Get("/", userHandler.Home)
	r.Get("/signup", userHandler.SignUp)
	r.Post("/users", userHandler.Create)
	r.Get("/signin", userHandler.SignIn)
	r.Post("/signin", userHandler.ProcessSignIn)
	r.Post("/signout", userHandler.ProcessSignOut)
	r.Get("/forgot-pw", userHandler.ForgotPassword)
	r.Post("/forgot-pw", userHandler.ProcessForgotPassword)
	r.Get("/reset-pw", userHandler.ResetPassword)
	r.Post("/reset-pw", userHandler.ProcessResetPassword)

	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", userHandler.CurrentUser)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	return r
}

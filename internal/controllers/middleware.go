package controllers

import (
	"log"
	"net/http"

	"github.com/samuelralmeida/product-catalog-api/internal/context"
	"github.com/samuelralmeida/product-catalog-api/internal/cookie"
)

type UserMiddleware struct {
	UserService UserService
}

func (umw UserMiddleware) SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := cookie.ReadCookie(r, cookie.CookieSession)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		user, err := umw.UserService.User(r.Context(), token)
		if err != nil {
			log.Println("user in set user middleware: %w", err)
			next.ServeHTTP(w, r)
			return
		}
		ctx := context.WithUser(r.Context(), user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// RequireUser requires "SetUser" middleware had been executed before it.
func (umw UserMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := context.User(r.Context())
		if user == nil {
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

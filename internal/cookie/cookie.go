package cookie

import (
	"fmt"
	"net/http"
)

// TODO: criar uma interface para injeção de depêndencia, facilitar testes

type cookieName string

const CookieSession cookieName = "session"

func NewCooke(name cookieName, value string) *http.Cookie {
	cookie := http.Cookie{
		Name:     string(name),
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	return &cookie
}

func SetCookie(w http.ResponseWriter, name cookieName, value string) {
	cookie := NewCooke(name, value)
	http.SetCookie(w, cookie)
}

func ReadCookie(r *http.Request, name cookieName) (string, error) {
	cookie, err := r.Cookie(string(name))
	if err != nil {
		return "", fmt.Errorf("%s: %w", name, err)
	}
	return cookie.Value, nil
}

func DeleteCookie(w http.ResponseWriter, name cookieName) {
	cookie := NewCooke(name, "")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}

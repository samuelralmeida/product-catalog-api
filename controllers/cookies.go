package controllers

import (
	"fmt"
	"net/http"
)

// TODO: criar uma interface para injeção de depêndencia, facilitar testes

type cookieName string

const cookieSession cookieName = "session"

func newCooke(name cookieName, value string) *http.Cookie {
	cookie := http.Cookie{
		Name:     string(name),
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	return &cookie
}

func setCookie(w http.ResponseWriter, name cookieName, value string) {
	cookie := newCooke(name, value)
	http.SetCookie(w, cookie)
}

func readCookie(r *http.Request, name cookieName) (string, error) {
	cookie, err := r.Cookie(string(cookieSession))
	if err != nil {
		return "", fmt.Errorf("%s: %w", name, err)
	}
	return cookie.Value, nil
}

func deleteCookie(w http.ResponseWriter, name cookieName) {
	cookie := newCooke(name, "")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

}

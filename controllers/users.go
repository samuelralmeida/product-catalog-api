package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/samuelralmeida/product-catalog-api/models"
)

type User struct {
	Templates struct {
		Signup Template
		Signin Template
	}
	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u User) SignUp(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email     string
		CSRFField template.HTML
	}
	data.Email = r.FormValue("email")
	data.CSRFField = csrf.TemplateField(r)
	u.Templates.Signup.Execute(w, r, data)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something wrong", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		// TODO: show warning about not being able to sign the user in
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	setCookie(w, cookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u User) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signin.Execute(w, r, data)
}

func (u User) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something wrong", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	setCookie(w, cookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u User) CurrentUser(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, cookieSession)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	user, err := u.SessionService.User(token)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Current user: %v\n", user)
}

func (u User) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, cookieSession)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	err = u.SessionService.Delete(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	deleteCookie(w, cookieSession)
	http.Redirect(w, r, "/signin", http.StatusFound)

}

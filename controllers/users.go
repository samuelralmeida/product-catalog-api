package controllers

import (
	"fmt"
	"net/http"

	"github.com/samuelralmeida/product-catalog-api/models"
)

type User struct {
	Templates struct {
		New    Template
		Signin Template
	}
	UserService *models.UserService
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
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
	fmt.Fprintf(w, "User created: %v", user)
}

func (u User) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signin.Execute(w, data)
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
	fmt.Fprintf(w, "User authenticate: %v", user)
}

package chi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samuelralmeida/product-catalog-api/internal/context"
	"github.com/samuelralmeida/product-catalog-api/internal/controllers"
	"github.com/samuelralmeida/product-catalog-api/internal/cookie"
)

type Users struct {
	Templates   controllers.HtmlTemplates
	UserService controllers.UserService
}

func (u Users) Home(w http.ResponseWriter, r *http.Request) {
	u.Templates.Signup.Execute(w, r, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	ctx := r.Context()

	user, session, err := u.UserService.Create(ctx, data.Email, data.Password)
	if err != nil {
		if user == nil {
			log.Println(fmt.Errorf("create user: %w", err))
			u.Templates.Signup.Execute(w, r, data, err)
			return
		}
		log.Println(fmt.Errorf("create session: %w", err))
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	cookie.SetCookie(w, cookie.CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) SignUp(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signup.Execute(w, r, data)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signin.Execute(w, r, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	ctx := r.Context()

	session, err := u.UserService.Autheticate(ctx, data.Email, data.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "error to authenticate user", http.StatusInternalServerError)
		return
	}

	cookie.SetCookie(w, cookie.CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	user := context.User(r.Context())
	fmt.Fprintf(w, "Current user: %v\n", user.Email)
}

func (u Users) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := cookie.ReadCookie(r, cookie.CookieSession)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	ctx := r.Context()

	err = u.UserService.SignOut(ctx, sessionToken)
	if err != nil {
		log.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	cookie.DeleteCookie(w, cookie.CookieSession)
	http.Redirect(w, r, "/signin", http.StatusFound)
}

func (u Users) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.ForgotPassword.Execute(w, r, data)
}

func (u Users) ProcessForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")

	ctx := r.Context()

	_, err := u.UserService.ForgotPassword(ctx, data.Email)
	if err != nil {
		// TODO: tratar se o email informado não existe
		log.Println(err)
		http.Error(w, "error to generate reset password", http.StatusInternalServerError)
		return
	}
	u.Templates.CheckYourEmail.Execute(w, r, data)
}

func (u Users) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token string
	}
	data.Token = r.FormValue("token")
	u.Templates.ResetPassword.Execute(w, r, data)
}

func (u Users) ProcessResetPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token    string
		Password string
	}
	data.Token = r.FormValue("token")
	data.Password = r.FormValue("password")

	ctx := r.Context()

	session, err := u.UserService.ResetPassword(ctx, data.Token, data.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "error to reset password", http.StatusInternalServerError)
		return
	}

	cookie.SetCookie(w, cookie.CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)

}

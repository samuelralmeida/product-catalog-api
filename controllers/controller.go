package controllers

import (
	"context"
	"net/http"

	"github.com/samuelralmeida/product-catalog-api/entity"
	"github.com/samuelralmeida/product-catalog-api/internal/env"
	"github.com/samuelralmeida/product-catalog-api/templates"
)

type UserService interface {
	Create(ctx context.Context, email string, password string) (*entity.User, *entity.Session, error)
	Autheticate(ctx context.Context, email string, password string) (*entity.Session, error)
	SignOut(ctx context.Context, sessionToken string) error
	ForgotPassword(ctx context.Context, email string) (*entity.PasswordReset, error)
	ResetPassword(ctx context.Context, resetPasswordToken string, password string) (*entity.Session, error)
	User(ctx context.Context, sessionToken string) (*entity.User, error)
}

type Controller struct {
	UserService UserService
	Config      *env.Config
}

type Template interface {
	Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error)
}

type HtmlTemplates struct {
	Home           Template
	Signup         Template
	Signin         Template
	ForgotPassword Template
	CheckYourEmail Template
	ResetPassword  Template
}

func MustParseTemplates() HtmlTemplates {
	return HtmlTemplates{
		Home:           templates.MustParseFS(templates.FS, "layout-page.gohtml", "home.gohtml"),
		Signup:         templates.MustParseFS(templates.FS, "layout-page.gohtml", "signup.gohtml"),
		Signin:         templates.MustParseFS(templates.FS, "layout-page.gohtml", "signin.gohtml"),
		ForgotPassword: templates.MustParseFS(templates.FS, "layout-page.gohtml", "forgot-pw.gohtml"),
		CheckYourEmail: templates.MustParseFS(templates.FS, "layout-page.gohtml", "check-your-email.gohtml"),
		ResetPassword:  templates.MustParseFS(templates.FS, "layout-page.gohtml", "reset-pw.gohtml"),
	}
}

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

type ProductService interface {
	Products(ctx context.Context) (*[]entity.Product, error)
	CreateProduct(ctx context.Context, product *entity.Product) error
	Product(ctx context.Context, id uint) (*entity.Product, error)
}

type MeasurementService interface {
	Measurements(ctx context.Context) (*[]entity.Measurement, error)
	CreateMeasurement(ctx context.Context, measurement *entity.Measurement) error
	Measurement(ctx context.Context, sybol string) (*entity.Measurement, error)
}

type ManufacturerService interface {
	Manufacturers(ctx context.Context) (*[]entity.Manufacturer, error)
	CreateManufacturer(ctx context.Context, manufactrer *entity.Manufacturer) error
	Manufacturer(ctx context.Context, id uint) (*entity.Manufacturer, error)
}

type Controller struct {
	UserService         UserService
	ProductService      ProductService
	MeasurementService  MeasurementService
	ManufacturerService ManufacturerService
	Config              *env.Config
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

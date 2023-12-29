package views

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/samuelralmeida/product-catalog-api/context"
	"github.com/samuelralmeida/product-catalog-api/models"
)

type public interface {
	Public() string
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func MustParseFS(fs fs.FS, pattern ...string) Template {
	tpl := template.New(pattern[0])

	// placeholder function to allow to parse templates in compilation time
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implemented")
			},
			"currentUser": func() (template.HTML, error) {
				return "", fmt.Errorf("currentUser not implemented")
			},
			"errors": func() []string {
				return []string{
					"Don't do that!",
					"The email address you provided is already associated with an account.",
					"Something went wrong.",
				}
			},
		},
	)

	tpl = template.Must(tpl.ParseFS(fs, pattern...))

	return Template{HtmlTemplate: tpl}
}

type Template struct {
	HtmlTemplate *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	// cloning avoid race condition during multiple requests
	tpl, err := t.HtmlTemplate.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "there was an error rendering the page", http.StatusInternalServerError)
		return
	}
	errMsgs := errMessages(errs...)

	// as funções abaixo são "closure" e podem ser chamadas mais de uma vez ao longo da renderização
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
			"errors": func() []string {
				return errMsgs
			},
		},
	)

	w.Header().Set("Context-Type", "text/html; charset=utf-8")

	// Implementing a buffer for rendering the entire page before writing it to the response.
	// Without buffering, pages with errors will render incompletely in the browser.
	// However, using a buffer for rendering extensive pages might potentially impact performance.
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "there was an error executing the template", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

func errMessages(errs ...error) []string {
	var msgs []string
	for _, err := range errs {
		var pubErr public
		if errors.As(err, &pubErr) {
			msgs = append(msgs, pubErr.Public())
		} else {
			fmt.Println(err)
			msgs = append(msgs, "something went wrong")
		}
	}
	return msgs
}

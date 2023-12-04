package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func MustParseFS(fs fs.FS, pattern ...string) Template {
	tpl := template.Must(template.ParseFS(fs, pattern...))
	return Template{HtmlTemplate: tpl}
}

func Parse(filepath string) (Template, error) {
	// TODO: can add template.Must() here and avoid Must func above
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{HtmlTemplate: tpl}, nil
}

type Template struct {
	HtmlTemplate *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Context-Type", "text/html; charset=utf-8")
	err := t.HtmlTemplate.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "there was an error executing the template", http.StatusInternalServerError)
		return
	}
}

package controllers

import (
	"net/http"

	"github.com/samuelralmeida/product-catalog-api/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func ProductHandler(tpl views.Template) http.HandlerFunc {
	products := []struct {
		Name string
		Gtin string
	}{
		{
			Name: "Dipirona",
			Gtin: "78945612587140369",
		},
		{
			Name: "Clonazepan",
			Gtin: "46584646586458654",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, products)
	}
}

package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func ProductHandler(tpl Template) http.HandlerFunc {
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
		tpl.Execute(w, r, products)
	}
}

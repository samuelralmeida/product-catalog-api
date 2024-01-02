package chi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/samuelralmeida/product-catalog-api/controllers"
)

type Products struct {
	ProductService controllers.ProducService
}

func (p Products) List(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	ctx := r.Context()
	products, err := p.ProductService.List(ctx)
	if err != nil {
		log.Println("list products: %w", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println("json products: %w", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

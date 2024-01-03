package chi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/samuelralmeida/product-catalog-api/controllers"
	"github.com/samuelralmeida/product-catalog-api/controllers/dto"
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

	render.JSON(w, r, products)
}

func (p Products) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	var requestBody dto.InsertProductRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("decode create product request:", err)
		http.Error(w, "internal error", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	product := requestBody.ToProduct()
	err = p.ProductService.Create(ctx, product)
	if err != nil {
		log.Println("create product:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, product)
}

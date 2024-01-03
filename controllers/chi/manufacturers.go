package chi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/samuelralmeida/product-catalog-api/controllers"
	"github.com/samuelralmeida/product-catalog-api/controllers/dto"
)

type Manufacturers struct {
	ManufacturerService controllers.ManufacturerService
}

func (m Manufacturers) List(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	ctx := r.Context()
	products, err := m.ManufacturerService.Manufacturers(ctx)
	if err != nil {
		log.Println("list manufacturers: %w", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, products)
}

func (m Manufacturers) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	var requestBody dto.InsertManufaturerRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("decode create manufacturers request:", err)
		http.Error(w, "bad request error", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	manufacturer := requestBody.ToEntity()
	err = m.ManufacturerService.CreateManufacturer(ctx, manufacturer)
	if err != nil {
		log.Println("create measurement:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, manufacturer)
}

func (m Manufacturers) Manufacturer(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	rawId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Println("parse product id:", err)
		http.Error(w, "bad request error", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	product, err := m.ManufacturerService.Manufacturer(ctx, uint(id))
	if err != nil {
		log.Println("get product by id:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, product)
}

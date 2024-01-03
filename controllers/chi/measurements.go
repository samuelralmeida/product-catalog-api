package chi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/samuelralmeida/product-catalog-api/controllers"
	"github.com/samuelralmeida/product-catalog-api/controllers/dto"
)

type Measurements struct {
	MeasurementService controllers.MeasurementService
}

func (m Measurements) List(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	ctx := r.Context()
	products, err := m.MeasurementService.Measurements(ctx)
	if err != nil {
		log.Println("list products: %w", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, products)
}

func (m Measurements) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	var requestBody dto.InsertMeasurementRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("decode create measurement request:", err)
		http.Error(w, "bad request error", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	measurement := requestBody.ToEntity()
	err = m.MeasurementService.CreateMeasurement(ctx, measurement)
	if err != nil {
		log.Println("create measurement:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, measurement)
}

func (m Measurements) Measurement(w http.ResponseWriter, r *http.Request) {
	// TODO: use go-playground/form to parse data request
	// TODO: use go-playground/validator to validate data

	symbol := chi.URLParam(r, "symbol")
	ctx := r.Context()
	product, err := m.MeasurementService.Measurement(ctx, symbol)
	if err != nil {
		log.Println("get measurement by symbol:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, product)
}

package dto

import "github.com/samuelralmeida/product-catalog-api/entity"

type InsertMeasurementRequest struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (r *InsertMeasurementRequest) ToEntity() *entity.Measurement {
	return &entity.Measurement{
		Name:   r.Name,
		Symbol: r.Symbol,
	}
}

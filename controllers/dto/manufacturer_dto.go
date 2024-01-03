package dto

import "github.com/samuelralmeida/product-catalog-api/entity"

type InsertManufaturerRequest struct {
	Name string `json:"name"`
}

func (r *InsertManufaturerRequest) ToEntity() *entity.Manufacturer {
	return &entity.Manufacturer{
		Name: r.Name,
	}
}

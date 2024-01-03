package dto

import "github.com/samuelralmeida/product-catalog-api/entity"

type InsertProductRequest struct {
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Presentation            string `json:"presentation"`
	HeightMM                *int   `json:"height_mm"`
	WidthMM                 *int   `json:"width_mm"`
	LengthMM                *int   `json:"length_mm"`
	Quantity                int    `json:"quantity"`
	UnitOfMeasurementSymbol string `json:"unit_measurement_symbol"`
	StorageCondition        string `json:"storage_condition"`
	GrossWeightG            *int   `json:"gross_weight_g"`
	NetWeightG              *int   `json:"net_weight_g"`
	Brand                   string `json:"brand"`
	Ncm                     string `json:"ncm"`
	Gtin                    string `json:"gtin"`
	ManufacturerId          uint   `json:"manufacturer_id"`
	GroupId                 uint   `json:"group_id"`
	AssociatedConditionID   *uint  `json:"associated_condition_id"`
	UmbrellaItemID          *uint  `json:"umbrella_item_id"`
}

func (r *InsertProductRequest) ToProduct() *entity.Product {
	return &entity.Product{
		Name:                    r.Name,
		Description:             r.Description,
		Presentation:            r.Presentation,
		HeightMM:                r.HeightMM,
		WidthMM:                 r.WidthMM,
		LengthMM:                r.LengthMM,
		Quantity:                r.Quantity,
		StorageCondition:        r.StorageCondition,
		GrossWeightG:            r.GrossWeightG,
		NetWeightG:              r.NetWeightG,
		Brand:                   r.Brand,
		Ncm:                     r.Ncm,
		Gtins:                   []string{r.Gtin},
		ManufacturerID:          r.ManufacturerId,
		GroupID:                 r.GroupId,
		AssociatedConditionID:   r.AssociatedConditionID,
		UnitOfMeasurementSymbol: r.UnitOfMeasurementSymbol,
		UmbrellaItemID:          r.UmbrellaItemID,
	}
}

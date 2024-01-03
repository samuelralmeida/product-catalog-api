package services

import (
	"context"

	"github.com/samuelralmeida/product-catalog-api/entity"
)

type ProductUseCases interface {
	List(ctx context.Context) (*[]entity.Product, error)
	Create(ctx context.Context, product *entity.Product) error
}

type ManufacturerUseCases interface {
	Manufacturer(ctx context.Context, id uint) (*entity.Manufacturer, error)
	Create(ctx context.Context, manufacturer *entity.Manufacturer) error
}

type MeasurementUseCases interface {
	Measurement(ctx context.Context, symbol string) (*entity.Measurement, error)
	Create(ctx context.Context, measurement *entity.Measurement) error
}

type ProductService struct {
	ProductUseCases      ProductUseCases
	ManufacturerUseCases ManufacturerUseCases
	MeasurementUseCases  MeasurementUseCases
}

func (ps *ProductService) List(ctx context.Context) (*[]entity.Product, error) {
	return ps.ProductUseCases.List(ctx)
}

func (ps *ProductService) Create(ctx context.Context, product *entity.Product) error {
	return ps.ProductUseCases.Create(ctx, product)
}

package services

import (
	"context"

	"github.com/samuelralmeida/product-catalog-api/entity"
)

type ProductUseCases interface {
	Products(ctx context.Context) (*[]entity.Product, error)
	Create(ctx context.Context, product *entity.Product) error
	Product(ctx context.Context, id uint) (*entity.Product, error)
}

type ManufacturerUseCases interface {
	Manufacturers(ctx context.Context) (*[]entity.Manufacturer, error)
	Manufacturer(ctx context.Context, id uint) (*entity.Manufacturer, error)
	Create(ctx context.Context, manufacturer *entity.Manufacturer) error
}

type MeasurementUseCases interface {
	Measurements(ctx context.Context) (*[]entity.Measurement, error)
	Measurement(ctx context.Context, symbol string) (*entity.Measurement, error)
	Create(ctx context.Context, measurement *entity.Measurement) error
}

type ProductService struct {
	ProductUseCases      ProductUseCases
	ManufacturerUseCases ManufacturerUseCases
	MeasurementUseCases  MeasurementUseCases
}

func (ps *ProductService) Products(ctx context.Context) (*[]entity.Product, error) {
	return ps.ProductUseCases.Products(ctx)
}

func (ps *ProductService) CreateProduct(ctx context.Context, product *entity.Product) error {
	return ps.ProductUseCases.Create(ctx, product)
}

func (ps *ProductService) Product(ctx context.Context, id uint) (*entity.Product, error) {
	return ps.ProductUseCases.Product(ctx, id)
}

func (ps *ProductService) Measurements(ctx context.Context) (*[]entity.Measurement, error) {
	return ps.MeasurementUseCases.Measurements(ctx)
}

func (ps *ProductService) CreateMeasurement(ctx context.Context, measurement *entity.Measurement) error {
	return ps.MeasurementUseCases.Create(ctx, measurement)
}

func (ps *ProductService) Measurement(ctx context.Context, symbol string) (*entity.Measurement, error) {
	return ps.MeasurementUseCases.Measurement(ctx, symbol)
}

func (ps *ProductService) Manufacturers(ctx context.Context) (*[]entity.Manufacturer, error) {
	return ps.ManufacturerUseCases.Manufacturers(ctx)
}

func (ps *ProductService) CreateManufacturer(ctx context.Context, manufacturer *entity.Manufacturer) error {
	return ps.ManufacturerUseCases.Create(ctx, manufacturer)
}

func (ps *ProductService) Manufacturer(ctx context.Context, id uint) (*entity.Manufacturer, error) {
	return ps.ManufacturerUseCases.Manufacturer(ctx, id)
}

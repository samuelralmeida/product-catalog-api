package services

import (
	"context"

	"github.com/samuelralmeida/product-catalog-api/entity"
)

type ProductUseCases interface {
	List(ctx context.Context) (*[]entity.Product, error)
}

type ProductService struct {
	ProductUseCases ProductUseCases
}

func (ps *ProductService) List(ctx context.Context) (*[]entity.Product, error) {
	return ps.ProductUseCases.List(ctx)
}

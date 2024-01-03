package product

import (
	"context"

	"github.com/samuelralmeida/product-catalog-api/entity"
)

type ProductRepository interface {
	Products(ctx context.Context) (*[]entity.Product, error)
	Create(ctx context.Context, product *entity.Product) error
}

type UseCase struct {
	Repository ProductRepository
}

func (us *UseCase) List(ctx context.Context) (*[]entity.Product, error) {
	return us.Repository.Products(ctx)
}

func (us *UseCase) Create(ctx context.Context, product *entity.Product) error {
	return us.Repository.Create(ctx, product)
}
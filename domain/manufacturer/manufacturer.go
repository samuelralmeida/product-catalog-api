package manufacturer

import (
	"context"

	"github.com/samuelralmeida/product-catalog-api/entity"
)

type ManufacturerRepository interface {
	Manufacturer(ctx context.Context, id uint) (*entity.Manufacturer, error)
	Create(ctx context.Context, manufacturer *entity.Manufacturer) error
}

type UseCases struct {
	Repository ManufacturerRepository
}

func (uc *UseCases) Create(ctx context.Context, manufacturer *entity.Manufacturer) error {
	return uc.Repository.Create(ctx, manufacturer)
}

func (uc *UseCases) Manufacturer(ctx context.Context, id uint) (*entity.Manufacturer, error) {
	return uc.Repository.Manufacturer(ctx, id)
}

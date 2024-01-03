package measurement

import (
	"context"

	"github.com/samuelralmeida/product-catalog-api/entity"
)

type MeasurementRepository interface {
	Measurement(ctx context.Context, symbol string) (*entity.Measurement, error)
	Create(ctx context.Context, measurement *entity.Measurement) error
}

type UseCases struct {
	Repository MeasurementRepository
}

func (uc *UseCases) Create(ctx context.Context, measurement *entity.Measurement) error {
	return uc.Repository.Create(ctx, measurement)
}

func (uc *UseCases) Measurement(ctx context.Context, symbol string) (*entity.Measurement, error) {
	return uc.Repository.Measurement(ctx, symbol)
}

package measurementpostgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/entity"
)

type MeasurementRepository struct {
	DB database.Database
}

const selectMeasurementBySymbolQuery = "select symbol, name from products.units_of_measurement where symbol = '$1'"

func (mr *MeasurementRepository) Measurement(ctx context.Context, symbol string) (*entity.Measurement, error) {
	var measurement entity.Measurement
	row := mr.DB.QueryRowContext(ctx, selectMeasurementBySymbolQuery, symbol)
	err := row.Scan(&measurement.Symbol, &measurement.Name)
	if err != nil {
		return nil, fmt.Errorf("select measurement by symbol: %w", err)
	}
	return &measurement, nil
}

const insertMeasurementQuery = "INSERT INTO products.units_of_measurement (symbol, name, deleted_at) VALUES($1, $2, $3) returning id;"

func (mr *MeasurementRepository) Create(ctx context.Context, measurement *entity.Measurement) error {
	deletedAt := sql.NullTime{}
	if measurement.DeletedAt != nil {
		deletedAt.Time = *measurement.DeletedAt
		deletedAt.Valid = true
	}

	_, err := mr.DB.ExecContext(ctx, insertMeasurementQuery, measurement.Symbol, measurement.Name, deletedAt)
	if err != nil {
		return fmt.Errorf("insert measurement: %w", err)
	}

	return nil
}

package manufacturerpostgres

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/entity"
)

type ManufacturerRepository struct {
	DB database.Database
}

const selectmanufacturersQuery = "select id, name from products.manufacturers"

func (pr *ManufacturerRepository) Manufacturers(ctx context.Context) (*[]entity.Manufacturer, error) {
	var manufacturers []entity.Manufacturer
	rows, err := pr.DB.QueryContext(ctx, selectmanufacturersQuery)
	if err != nil {
		return nil, fmt.Errorf("select manufacturers: %w", err)
	}

	for rows.Next() {
		var manufacturer entity.Manufacturer
		err := rows.Scan(&manufacturer.ID, &manufacturer.Name)
		if err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}
		manufacturers = append(manufacturers, manufacturer)
	}

	return &manufacturers, nil
}

const selectManufacturerByIdQuery = "select id, name from products.manufacturers where id = $1"

func (mr *ManufacturerRepository) Manufacturer(ctx context.Context, id uint) (*entity.Manufacturer, error) {
	var manufacturer entity.Manufacturer
	row := mr.DB.QueryRowContext(ctx, selectManufacturerByIdQuery, id)
	err := row.Scan(&manufacturer.ID, &manufacturer.Name)
	if err != nil {
		return nil, fmt.Errorf("select manufactrer by id: %w", err)
	}
	return &manufacturer, nil
}

const insertManufacturerQuery = "INSERT INTO products.manufacturers (name, deleted_at) VALUES($1, $2) returning id;"

func (mr *ManufacturerRepository) Create(ctx context.Context, manufacturer *entity.Manufacturer) error {
	err := mr.DB.QueryRowContext(ctx, insertManufacturerQuery, manufacturer.Name, manufacturer.DeletedAt).Scan(&manufacturer.ID)
	if err != nil {
		return fmt.Errorf("insert manufacturer: %w", err)
	}

	return nil
}

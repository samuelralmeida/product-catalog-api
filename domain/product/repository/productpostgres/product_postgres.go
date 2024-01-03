package productpostgres

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/product-catalog-api/database"
	"github.com/samuelralmeida/product-catalog-api/entity"
)

type ProductRepository struct {
	DB database.Database
}

const selectProductsQuery = "select id, name, description, presentation from products.products"

func (pr *ProductRepository) Products(ctx context.Context) (*[]entity.Product, error) {
	var products []entity.Product
	rows, err := pr.DB.QueryContext(ctx, selectProductsQuery)
	if err != nil {
		return nil, fmt.Errorf("select products: %w", err)
	}

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Presentation)
		if err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}
		products = append(products, product)
	}

	return &products, nil
}

const insertProductQuery = `
	INSERT INTO products.products (
		gross_weight_g, height_mm, length_mm, net_weight_g, quantity, width_mm, manufacturer_id,
		brand, description, name, ncm, presentation, storage_condition,
		unit_of_measurement_symbol, group_id, umbrella_item_id
	)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	RETURNING id;
`

func (pr *ProductRepository) Create(ctx context.Context, product *entity.Product) error {
	err := pr.DB.QueryRowContext(ctx, insertProductQuery,
		product.GrossWeightG, product.HeightMM, product.LengthMM, product.NetWeightG, product.Quantity, product.WidthMM,
		product.ManufacturerID, product.Brand, product.Description, product.Name, product.Ncm, product.Presentation,
		product.StorageCondition, product.UnitOfMeasurementSymbol, product.GroupID, product.UmbrellaItemID,
	).Scan(&product.ID)

	if err != nil {
		return fmt.Errorf("insert product: %w", err)
	}

	return nil
}

const selectProductByIdQuery = `
	SELECT
		id, gross_weight_g, height_mm, length_mm, net_weight_g, quantity,
		width_mm, manufacturer_id, brand, description, name, ncm, presentation,
		storage_condition, unit_of_measurement_symbol, group_id, umbrella_item_id
	FROM products.products
	WHERE id = $1
`

func (pr *ProductRepository) Product(ctx context.Context, id uint) (*entity.Product, error) {
	var product entity.Product
	err := pr.DB.QueryRowContext(ctx, selectProductByIdQuery, id).Scan(
		&product.ID, &product.GrossWeightG, &product.HeightMM, &product.LengthMM, &product.NetWeightG, &product.Quantity,
		&product.WidthMM, &product.ManufacturerID, &product.Brand, &product.Description, &product.Name, &product.Ncm,
		&product.Presentation, &product.StorageCondition, &product.UnitOfMeasurementSymbol, &product.GroupID, &product.UmbrellaItemID,
	)

	if err != nil {
		return nil, fmt.Errorf("get product by id: %w", err)
	}

	return &product, nil
}

/*
func (p *ProductPostgresql) ProductByGtin(gtin string) (*domain.Product, error) {
	var prod domain.Product
	err := p.db.Where("? = ANY(gtins)", gtin).First(&prod).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.ErrProductNotFound
	}
	return &prod, exception.NewInternalError(err)
}

func (p *ProductPostgresql) Create(prod *domain.Product) error {
	return exception.NewInternalError(p.db.Create(prod).Error)
}

func (p *ProductPostgresql) List(pagination domain.Pagination) (*[]domain.Product, error) {
	var products []domain.Product
	query := p.db.Offset(pagination.Offset()).Limit(pagination.PageSize())
	for _, order := range pagination.OrderBy() {
		query = query.Order(order)
	}
	err := query.Find(&products).Error
	return &products, exception.NewInternalError(err)
}
*/

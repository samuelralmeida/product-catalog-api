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

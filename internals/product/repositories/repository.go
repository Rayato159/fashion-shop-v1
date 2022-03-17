package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/fashion-shop/v1/internals/models"
	"github.com/rayato159/fashion-shop/v1/internals/product"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) product.Repository {
	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(p *models.CreateProduct) error {
	createProductQuery := fmt.Sprintf(`
		INSERT INTO products (category_id) 
		VALUES((SELECT category_id FROM categories WHERE category_id = %d))
	`, p.CategoryId)

	_, err := r.db.Exec(createProductQuery)
	if err != nil {
		return fmt.Errorf("query has failed with error: %w", err)
	}
	return nil
}

func (r *productRepo) GetAllProduct(f *models.ProductFilter) ([]*models.Product, error) {
	getProductQuery := `
	SELECT 
		products.product_id, 
		products.category_id,
		categories.gender,
		categories.size,
		categories.price
	FROM products
		LEFT JOIN categories
			ON products.category_id = categories.category_id
	`

	if f != nil {
		var filters = make([]string, 0)

		if f.CategoryId != 0 {
			filters = append(filters, fmt.Sprintf(`products.category_id = %d`, f.CategoryId))
		}
		if f.Gender != "" {
			filters = append(filters, fmt.Sprintf(`categories.gender = '%s'`, f.Gender))
		}
		if f.Size != "" {
			filters = append(filters, fmt.Sprintf(`categories.size = '%s'`, f.Size))
		}

		for i, items := range filters {
			if i == 0 {
				getProductQuery += "\nWHERE"
			} else {
				getProductQuery += "\nAND"
			}
			getProductQuery += fmt.Sprintf("\n%s", items)
		}
	}

	rows, err := r.db.Queryx(getProductQuery)
	if err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}
	defer rows.Close()

	var products = make([]*models.Product, 0)
	for rows.Next() {
		var product models.Product
		if err := rows.StructScan(&product); err != nil {
			return nil, fmt.Errorf("query has failed with error: %w", err)
		}
		products = append(products, &product)
	}
	return products, nil
}

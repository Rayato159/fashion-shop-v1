package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/fashion-shop/v1/internals/category"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) category.Repository {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) CreateCategory(c *models.CreateCategory) error {
	createCategoryQuery := fmt.Sprintf(`
	INSERT INTO categories (gender, size, price, color_id, pattern_id, figure_id) VALUES
	(
		'%s',
		'%s',
		%.2f,
		(SELECT color_id FROM colors WHERE color = '%s' LIMIT 1),
		(SELECT pattern_id FROM patterns WHERE pattern = '%s' LIMIT 1),
		(SELECT figure_id FROM figures WHERE figure = '%s' LIMIT 1)
	);`, c.Gender, c.Size, c.Price, c.Color, c.Pattern, c.Figure)

	_, err := r.db.Exec(createCategoryQuery)
	if err != nil {
		return fmt.Errorf("query has failed with error: %w", err)
	}
	return nil
}

func (r *categoryRepo) GetAllCategory() ([]*models.Category, error) {
	getAllCategoryQuery := `SELECT * FROM categories;`

	rows, err := r.db.Queryx(getAllCategoryQuery)
	if err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}

	var categories = make([]*models.Category, 0)
	for rows.Next() {
		var category models.Category
		if err := rows.StructScan(&category); err != nil {
			return nil, fmt.Errorf("query has failed with error: %w", err)
		}
		categories = append(categories, &category)
	}
	defer rows.Close()

	return categories, nil
}

package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/fashion-shop/v1/internals/color"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type colorRepo struct {
	db *sqlx.DB
}

func NewColorRepository(db *sqlx.DB) color.Repository {
	return &colorRepo{db: db}
}

func (r *colorRepo) GetAllColor(f *models.ColorFilter) ([]*models.Color, error) {
	colorQuery := `SELECT * FROM colors ORDER BY color DESC;`

	if f != nil {
		colorQuery = fmt.Sprintf("SELECT * FROM colors WHERE color LIKE '%%%s%%' ORDER BY color DESC;", f.Color)
	}

	rows, err := r.db.Queryx(colorQuery)
	if err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}
	defer rows.Close()

	var colorLists = make([]*models.Color, 0)
	for rows.Next() {
		var color models.Color
		if err := rows.StructScan(&color); err != nil {
			return nil, fmt.Errorf("query has failed with error: %w", err)
		}
		colorLists = append(colorLists, &color)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}
	return colorLists, nil
}

func (r *colorRepo) GetColorByKey(f string) (*models.Color, error) {
	colorQuery := fmt.Sprintf("SELECT * FROM colors WHERE color = '%s';", f)

	rows, err := r.db.Queryx(colorQuery)
	if err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}
	defer rows.Close()

	colorOne := new(models.Color)
	for rows.Next() {
		if err := rows.StructScan(&colorOne); err != nil {
			return nil, fmt.Errorf("query has failed with error: %w", err)
		}
	}

	return colorOne, nil
}

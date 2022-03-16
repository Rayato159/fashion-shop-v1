package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/fashion-shop/v1/internals/models"
	"github.com/rayato159/fashion-shop/v1/internals/pattern"
)

type patternRepo struct {
	db *sqlx.DB
}

func NewPatternRepository(db *sqlx.DB) pattern.Repository {
	return &patternRepo{db: db}
}

func (r *patternRepo) GetAllPattern(f *models.FilterPattern) ([]*models.Pattern, error) {
	var getAllPatternQuery string
	if f != nil {
		getAllPatternQuery = fmt.Sprintf(`SELECT * FROM patterns WHERE pattern LIKE '%%%s%%'`, f.Pattern)
	} else {
		getAllPatternQuery = `SELECT * FROM patterns`
	}

	rows, err := r.db.Queryx(getAllPatternQuery)
	if err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}

	var patternLists = make([]*models.Pattern, 0)
	for rows.Next() {
		var pattern models.Pattern
		if err := rows.StructScan(&pattern); err != nil {
			return nil, fmt.Errorf("query has failed with error: %w", err)
		}
		patternLists = append(patternLists, &pattern)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}
	return patternLists, nil
}

func (r *patternRepo) GetPatternByKey(k string) (*models.Pattern, error) {
	patternQuery := fmt.Sprintf(`SELECT * FROM patterns WHERE pattern = '%s' LIMIT 1;`, k)

	pattern := new(models.Pattern)
	if err := r.db.Get(pattern, patternQuery); err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}
	return pattern, nil
}

func (r *patternRepo) CreatePattern(p *models.CreatePattern) error {
	createPatternQuery := fmt.Sprintf(`INSERT INTO patterns (pattern) VALUES('%s');`, p.Pattern)

	_, err := r.db.Exec(createPatternQuery)
	if err != nil {
		return fmt.Errorf("query has failed with error: %w", err)
	}
	return nil
}

func (r *patternRepo) CreatePatternButBulk(p []models.CreatePattern) error {
	createPatternBulkQuery := `
		INSERT INTO patterns (pattern)
		VALUES (:pattern)
	`
	_, err := r.db.NamedExec(createPatternBulkQuery, p)
	if err != nil {
		return fmt.Errorf("query has failed with error: %w", err)
	}
	return nil
}

func (r *patternRepo) DeletePattern(k string) error {
	deletePatternQuery := fmt.Sprintf(`DELETE FROM patterns WHERE pattern = '%s'`, k)

	_, err := r.db.Exec(deletePatternQuery)
	if err != nil {
		return fmt.Errorf("query has failed with error: %w", err)
	}
	return nil
}

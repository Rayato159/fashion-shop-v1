package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/fashion-shop/v1/internals/figure"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type figureRepo struct {
	db *sqlx.DB
}

func NewFigureRepository(db *sqlx.DB) figure.Repository {
	return &figureRepo{db: db}
}

func (r *figureRepo) GetAllFigure(f *models.FigureFilter) ([]*models.Figure, error) {
	var figureQuery string
	if f != nil {
		figureQuery = fmt.Sprintf(`SELECT * FROM figures WHERE figure LIKE '%%%s%%';`, f.Figure)
	} else {
		figureQuery = `SELECT * FROM figures;`
	}

	rows, err := r.db.Queryx(figureQuery)
	if err != nil {
		return nil, fmt.Errorf("with error: %w", err)
	}

	var figures = make([]*models.Figure, 0)
	for rows.Next() {
		var figure models.Figure
		if err := rows.StructScan(&figure); err != nil {
			return nil, fmt.Errorf("with error: %w", err)
		}
		figures = append(figures, &figure)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}
	return figures, nil
}

func (r *figureRepo) GetFigureByKey(k string) (*models.Figure, error) {
	figureQuery := fmt.Sprintf(`SELECT * FROM figures WHERE figure = '%s'`, k)

	figure := new(models.Figure)
	err := r.db.Get(figure, figureQuery)
	if err != nil {
		return nil, fmt.Errorf("query has failed with error: %w", err)
	}
	return figure, nil
}

func (r *figureRepo) CreateFigure(f *models.CreateFigure) error {
	createFigureQuery := fmt.Sprintf(`INSERT INTO figures (figure) VALUES('%s')`, f.Figure)

	_, err := r.db.Exec(createFigureQuery)
	if err != nil {
		return fmt.Errorf("query has failed with error: %w", err)
	}

	return nil
}

func (r *figureRepo) CreateFigureButBulk(f []models.CreateFigure) error {
	createFigureQuery := `
		INSERT INTO figures (figure)
		VALUES(:figure)
	`

	_, err := r.db.NamedExec(createFigureQuery, f)
	if err != nil {
		return fmt.Errorf("query has failed with error: %w", err)
	}
	return nil
}

func (r *figureRepo) DeleteFigure(k string) error {
	deleteFigureQuery := fmt.Sprintf(`DELETE FROM figures WHERE figure = '%s'`, k)

	_, err := r.db.Exec(deleteFigureQuery)
	if err != nil {
		return fmt.Errorf("query has failed with error: %w", err)
	}
	return nil
}

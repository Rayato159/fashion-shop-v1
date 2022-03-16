package figure

import "github.com/rayato159/fashion-shop/v1/internals/models"

type Repository interface {
	GetAllFigure(f *models.FigureFilter) ([]*models.Figure, error)
	GetFigureByKey(k string) (*models.Figure, error)
	CreateFigure(f *models.CreateFigure) error
	CreateFigureButBulk(f []models.CreateFigure) error
	DeleteFigure(k string) error
}

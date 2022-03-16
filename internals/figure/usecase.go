package figure

import "github.com/rayato159/fashion-shop/v1/internals/models"

type UseCase interface {
	GetAllFigure(f *models.FigureFilter) ([]*models.Figure, error)
	GetFigureByKey(k string) (*models.Figure, error)
	CreateFigure(f *models.CreateFigure) error
	CreateFigureButBulk(f []models.CreateFigure) ([]*models.Figure, error)
	DeleteFigure(k string) error
}

package color

import (
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type UseCase interface {
	GetAllColor(f *models.ColorFilter) ([]*models.Color, error)
	GetColorByKey(f string) (*models.Color, error)
	CreateColor(c *models.CreateColor) (*models.Color, error)
	CreateColorButBulk(c []models.CreateColor) ([]*models.Color, error)
	DeleteColor(c string) error
}

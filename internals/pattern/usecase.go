package pattern

import "github.com/rayato159/fashion-shop/v1/internals/models"

type UseCase interface {
	GetAllPattern(f *models.FilterPattern) ([]*models.Pattern, error)
	GetPatternByKey(k string) (*models.Pattern, error)
	CreatePattern(p *models.CreatePattern) error
	CreatePatternButBulk(p []models.CreatePattern) ([]*models.Pattern, error)
	DeletePattern(k string) error
}

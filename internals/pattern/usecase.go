package pattern

import "github.com/rayato159/fashion-shop/v1/internals/models"

type UseCase interface {
	GetAllPattern(f *models.FilterPattern) (*models.Pattern, error)
}

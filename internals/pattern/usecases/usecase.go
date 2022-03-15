package usecases

import (
	"github.com/rayato159/fashion-shop/v1/internals/models"
	"github.com/rayato159/fashion-shop/v1/internals/pattern"
)

type patternUC struct {
	patternRepo pattern.Repository
}

func NewPatternUseCase(patternRepo pattern.Repository) pattern.UseCase {
	return &patternUC{patternRepo: patternRepo}
}

func (u *patternUC) GetAllPattern(f *models.FilterPattern) (*models.Pattern, error) {
	return nil, nil
}

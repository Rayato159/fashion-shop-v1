package usecases

import (
	"fmt"

	"github.com/rayato159/fashion-shop/v1/internals/models"
	"github.com/rayato159/fashion-shop/v1/internals/pattern"
)

type patternUC struct {
	patternRepo pattern.Repository
}

func NewPatternUseCase(patternRepo pattern.Repository) pattern.UseCase {
	return &patternUC{patternRepo: patternRepo}
}

func (u *patternUC) GetAllPattern(f *models.FilterPattern) ([]*models.Pattern, error) {
	patterns, err := u.patternRepo.GetAllPattern(f)
	if err != nil {
		return nil, fmt.Errorf("error, get patterns has failed with error: %w", err)
	}
	return patterns, nil
}

func (u *patternUC) GetPatternByKey(k string) (*models.Pattern, error) {
	pattern, err := u.patternRepo.GetPatternByKey(k)
	if err != nil {
		return nil, fmt.Errorf("error, get pattern by key has failed with error: %w", err)
	}
	return pattern, nil
}

func (u *patternUC) CreatePattern(p *models.CreatePattern) error {
	if err := u.patternRepo.CreatePattern(p); err != nil {
		return fmt.Errorf("error, create pattern has failed with error: %w", err)
	}
	return nil
}

func (u *patternUC) CreatePatternButBulk(p []models.CreatePattern) ([]*models.Pattern, error) {
	if err := u.patternRepo.CreatePatternButBulk(p); err != nil {
		return nil, fmt.Errorf("error, create patterns has failed with error: %w", err)
	}

	var patterns = make([]*models.Pattern, 0)
	for i := range p {
		pattern, err := u.patternRepo.GetPatternByKey(p[i].Pattern)
		if err != nil {
			return nil, fmt.Errorf("error, create patterns has failed with error: %w", err)
		}
		patterns = append(patterns, pattern)
	}
	return patterns, nil
}

func (u *patternUC) DeletePattern(k string) error {
	if err := u.patternRepo.DeletePattern(k); err != nil {
		return fmt.Errorf("error, delete patterns has failed with error: %w", err)
	}
	return nil
}

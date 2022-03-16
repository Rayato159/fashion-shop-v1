package usecases

import (
	"fmt"

	"github.com/rayato159/fashion-shop/v1/internals/color"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type colorUC struct {
	colorRepo color.Repository
}

func NewColorUseCase(colorRepo color.Repository) color.UseCase {
	return &colorUC{colorRepo: colorRepo}
}

func (u *colorUC) GetAllColor(f *models.ColorFilter) ([]*models.Color, error) {
	var colorLists []*models.Color

	dbData, err := u.colorRepo.GetAllColor(f)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	colorLists = dbData
	return colorLists, nil
}

func (u *colorUC) GetColorByKey(f string) (*models.Color, error) {
	colorOne, err := u.colorRepo.GetColorByKey(f)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return colorOne, nil
}

func (u *colorUC) CreateColor(c *models.CreateColor) (*models.Color, error) {
	err := u.colorRepo.CreateColor(c)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	colorResult, err := u.colorRepo.GetColorByKey(c.Color)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return colorResult, nil
}

func (u *colorUC) CreateColorButBulk(c []models.CreateColor) ([]*models.Color, error) {
	err := u.colorRepo.CreateColorButBulk(c)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	var colorResults = make([]*models.Color, 0)
	for i := range c {
		color, err := u.colorRepo.GetColorByKey(c[i].Color)
		if err != nil {
			return nil, fmt.Errorf("error: %w", err)
		}
		colorResults = append(colorResults, color)
	}
	return colorResults, nil
}

func (u *colorUC) DeleteColor(c string) error {
	err := u.colorRepo.DeleteColor(c)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

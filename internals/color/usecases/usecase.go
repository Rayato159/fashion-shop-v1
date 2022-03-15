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
	colorOne := new(models.Color)

	dbData, err := u.colorRepo.GetColorByKey(f)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	colorOne = dbData
	return colorOne, nil
}

package usecases

import (
	"fmt"

	"github.com/rayato159/fashion-shop/v1/internals/figure"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type figureUC struct {
	figureRepo figure.Repository
}

func NewFigureUseCase(figureRepo figure.Repository) figure.UseCase {
	return &figureUC{figureRepo: figureRepo}
}

func (u *figureUC) GetAllFigure(f *models.FigureFilter) ([]*models.Figure, error) {
	figures, err := u.figureRepo.GetAllFigure(f)
	if err != nil {
		return nil, fmt.Errorf("error, cant't get figures with error: %w", err)
	}
	return figures, nil
}

func (u *figureUC) GetFigureByKey(k string) (*models.Figure, error) {
	figure, err := u.figureRepo.GetFigureByKey(k)
	if err != nil {
		return nil, fmt.Errorf("error, cant't get figure with error: %w", err)
	}
	return figure, nil
}

func (u *figureUC) CreateFigure(f *models.CreateFigure) error {
	if err := u.figureRepo.CreateFigure(f); err != nil {
		return fmt.Errorf("error, cant't get figure with error: %w", err)
	}
	return nil
}

func (u *figureUC) CreateFigureButBulk(f []models.CreateFigure) ([]*models.Figure, error) {
	if err := u.figureRepo.CreateFigureButBulk(f); err != nil {
		return nil, fmt.Errorf("error, cant't create figures with error: %w", err)
	}

	var figures = make([]*models.Figure, 0)
	for i := range f {
		figure, err := u.figureRepo.GetFigureByKey(f[i].Figure)
		if err != nil {
			return nil, fmt.Errorf("error, cant't create figures with error: %w", err)
		}
		figures = append(figures, figure)
	}
	return figures, nil
}

func (u *figureUC) DeleteFigure(k string) error {
	if err := u.figureRepo.DeleteFigure(k); err != nil {
		return fmt.Errorf("error, cant't delete figures with error: %w", err)
	}
	return nil
}

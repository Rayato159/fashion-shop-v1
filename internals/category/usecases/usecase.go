package usecases

import (
	"fmt"

	"github.com/rayato159/fashion-shop/v1/internals/category"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type categoryUC struct {
	categoryRepo category.Repository
}

func NewCategoryUseCase(categoryRepo category.Repository) category.UseCase {
	return &categoryUC{categoryRepo: categoryRepo}
}

func (u *categoryUC) CreateCategory(c *models.CreateCategory) error {
	if err := u.categoryRepo.CreateCategory(c); err != nil {
		return fmt.Errorf("error, can't create category with error: %w", err)
	}
	return nil
}

func (u *categoryUC) GetAllCategory() ([]*models.Category, error) {
	categories, err := u.categoryRepo.GetAllCategory()
	if err != nil {
		return nil, fmt.Errorf("error, can't create category with error: %w", err)
	}
	return categories, nil
}

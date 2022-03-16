package category

import "github.com/rayato159/fashion-shop/v1/internals/models"

type Repository interface {
	CreateCategory(c *models.CreateCategory) error
	GetAllCategory() ([]*models.Category, error)
}

package product

import "github.com/rayato159/fashion-shop/v1/internals/models"

type UseCase interface {
	CreateProduct(p *models.CreateProduct) error
	GetAllProduct(f *models.ProductFilter) ([]*models.Product, error)
}

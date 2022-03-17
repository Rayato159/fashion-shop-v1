package usecases

import (
	"fmt"

	"github.com/rayato159/fashion-shop/v1/internals/models"
	"github.com/rayato159/fashion-shop/v1/internals/product"
)

type productUC struct {
	productRepo product.Repository
}

func NewProductUseCase(productRepo product.Repository) product.UseCase {
	return &productUC{productRepo: productRepo}
}

func (u *productUC) CreateProduct(p *models.CreateProduct) error {
	if err := u.productRepo.CreateProduct(p); err != nil {
		return fmt.Errorf("error, can't create product with error: %w", err)
	}
	return nil
}

func (u *productUC) GetAllProduct(f *models.ProductFilter) ([]*models.Product, error) {
	products, err := u.productRepo.GetAllProduct(f)
	if err != nil {
		return nil, fmt.Errorf("error, can't get product with error: %w", err)
	}
	return products, nil
}

package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/models"
	"github.com/rayato159/fashion-shop/v1/internals/product"
)

type productHandler struct {
	productUC product.UseCase
}

func NewProductHandler(productUC product.UseCase) product.Handler {
	return &productHandler{productUC: productUC}
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	np := new(models.CreateProduct)

	if err := c.BodyParser(np); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	fmt.Printf("%T\n", np.CategoryId)

	if err := h.productUC.CreateProduct(np); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Created",
	})
}

func (h *productHandler) GetAllProduct(c *fiber.Ctx) error {
	f := new(models.ProductFilter)

	if err := c.QueryParser(f); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	products, err := h.productUC.GetAllProduct(f)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "OK",
		"data":    products,
	})
}

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/category"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type categoryHandler struct {
	categoryUC category.UseCase
}

func NewCategoryHandler(categoryUC category.UseCase) category.Handler {
	return &categoryHandler{categoryUC: categoryUC}
}

func (h *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	cc := new(models.CreateCategory)

	if err := c.BodyParser(cc); err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	if err := h.categoryUC.CreateCategory(cc); err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	return c.Status(201).JSON(fiber.Map{
		"status":   fiber.StatusCreated,
		"messsage": "Created",
	})
}

func (h *categoryHandler) GetAllCategory(c *fiber.Ctx) error {
	categories, err := h.categoryUC.GetAllCategory()
	if err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	if len(categories) == 0 {
		return c.Status(404).JSON(fiber.ErrNotFound)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     categories,
	})
}

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/color"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type colorHandler struct {
	colorUC color.UseCase
}

func NewColorHandler(colorUC color.UseCase) color.Handler {
	return &colorHandler{colorUC: colorUC}
}

func (h *colorHandler) GetAllColor(c *fiber.Ctx) error {
	f := new(models.ColorFilter)
	if err := c.QueryParser(f); err != nil {
		return c.JSON(fiber.ErrInternalServerError)
	}

	colors, err := h.colorUC.GetAllColor(f)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if len(colors) == 0 {
		return c.Status(404).JSON(fiber.ErrNotFound)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     colors,
	})
}

func (h *colorHandler) GetColorByKey(c *fiber.Ctx) error {
	f := c.Params("color")

	colorData, err := h.colorUC.GetColorByKey(f)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     colorData,
	})
}

func (h *colorHandler) CreateColor(c *fiber.Ctx) error {
	newColor := new(models.CreateColor)
	if err := c.BodyParser(newColor); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	colorResult, err := h.colorUC.CreateColor(newColor)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":   fiber.StatusCreated,
		"messsage": "Created",
		"data":     colorResult,
	})
}

func (h *colorHandler) CreateColorButBulk(c *fiber.Ctx) error {
	newColor := new([]models.CreateColor)
	if err := c.BodyParser(&newColor); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	dbData, err := h.colorUC.CreateColorButBulk(*newColor)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"status":   fiber.StatusCreated,
		"messsage": "Created",
		"data":     dbData,
	})
}

func (h *colorHandler) DeleteColor(c *fiber.Ctx) error {
	colorDeleted := c.Params("color")

	err := h.colorUC.DeleteColor(colorDeleted)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
	})
}

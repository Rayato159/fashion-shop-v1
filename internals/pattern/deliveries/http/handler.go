package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/models"
	"github.com/rayato159/fashion-shop/v1/internals/pattern"
)

type patternHandler struct {
	patternUC pattern.UseCase
}

func NewPatternHandler(patternUC pattern.UseCase) pattern.Handler {
	return &patternHandler{patternUC: patternUC}
}

func (h *patternHandler) GetAllPattern(c *fiber.Ctx) error {
	q := new(models.FilterPattern)
	if err := c.QueryParser(q); err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	patterns, err := h.patternUC.GetAllPattern(q)
	if err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	} else if len(patterns) == 0 {
		return c.Status(404).JSON(fiber.ErrNotFound)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     patterns,
	})
}

func (h *patternHandler) GetPatternByKey(c *fiber.Ctx) error {
	key := c.Params("pattern")

	pattern, err := h.patternUC.GetPatternByKey(key)
	if err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     pattern,
	})
}

func (h *patternHandler) CreatePattern(c *fiber.Ctx) error {
	p := new(models.CreatePattern)

	if err := c.BodyParser(p); err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	err := h.patternUC.CreatePattern(p)
	if err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	return c.Status(201).JSON(fiber.Map{
		"status":   fiber.StatusCreated,
		"messsage": "Created",
	})
}

func (h *patternHandler) CreatePatternButBulk(c *fiber.Ctx) error {
	p := new([]models.CreatePattern)
	if err := c.BodyParser(p); err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	patterns, err := h.patternUC.CreatePatternButBulk(*p)
	if err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}

	return c.Status(201).JSON(fiber.Map{
		"status":   fiber.StatusCreated,
		"messsage": "Created",
		"data":     patterns,
	})
}

func (h *patternHandler) DeletePattern(c *fiber.Ctx) error {
	k := c.Params("pattern")

	if err := h.patternUC.DeletePattern(k); err != nil {
		return c.Status(500).JSON(fiber.ErrInternalServerError)
	}
	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
	})
}

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/figure"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

type figureHandler struct {
	figureUC figure.UseCase
}

func NewFigureHandler(figureUC figure.UseCase) figure.Handler {
	return &figureHandler{figureUC: figureUC}
}

func (h *figureHandler) GetAllFigure(c *fiber.Ctx) error {
	f := new(models.FigureFilter)
	if err := c.QueryParser(f); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	figures, err := h.figureUC.GetAllFigure(f)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if len(figures) == 0 {
		return c.Status(404).JSON(fiber.ErrNotFound)
	}
	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     figures,
	})
}

func (h *figureHandler) GetFigureByKey(c *fiber.Ctx) error {
	k := c.Params("figure")
	figure, err := h.figureUC.GetFigureByKey(k)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     figure,
	})
}

func (h *figureHandler) CreateFigure(c *fiber.Ctx) error {
	f := new(models.CreateFigure)

	if err := c.BodyParser(f); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if err := h.figureUC.CreateFigure(f); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":   fiber.StatusCreated,
		"messsage": "Created",
	})
}

func (h *figureHandler) CreateFigureButBulk(c *fiber.Ctx) error {
	f := new([]models.CreateFigure)

	if err := c.BodyParser(f); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	figures, err := h.figureUC.CreateFigureButBulk(*f)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"status":   fiber.StatusCreated,
		"messsage": "Created",
		"data":     figures,
	})
}

func (h *figureHandler) DeleteFigure(c *fiber.Ctx) error {
	k := c.Params("figure")

	if err := h.figureUC.DeleteFigure(k); err != nil {
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

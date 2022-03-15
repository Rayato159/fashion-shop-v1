package http

import (
	"fmt"

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
		return err
	}

	colors, err := h.colorUC.GetAllColor(f)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     colors,
	})
}

func (h *colorHandler) GetColorByKey(c *fiber.Ctx) error {
	f := c.Params("color")
	fmt.Println(f)

	colorData, err := h.colorUC.GetColorByKey(f)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"messsage": "OK",
		"data":     colorData,
	})
}

package figure

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetAllFigure(c *fiber.Ctx) error
	GetFigureByKey(c *fiber.Ctx) error
	CreateFigure(c *fiber.Ctx) error
	CreateFigureButBulk(c *fiber.Ctx) error
	DeleteFigure(c *fiber.Ctx) error
}

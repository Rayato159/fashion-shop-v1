package color

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetAllColor(c *fiber.Ctx) error
	GetColorByKey(c *fiber.Ctx) error
}

package pattern

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetAllPattern(c *fiber.Ctx) error
}

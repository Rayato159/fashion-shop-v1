package pattern

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetAllPattern(c *fiber.Ctx) error
	GetPatternByKey(c *fiber.Ctx) error
	CreatePattern(c *fiber.Ctx) error
	CreatePatternButBulk(c *fiber.Ctx) error
	DeletePattern(c *fiber.Ctx) error
}

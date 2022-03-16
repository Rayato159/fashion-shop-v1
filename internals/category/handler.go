package category

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateCategory(c *fiber.Ctx) error
	GetAllCategory(c *fiber.Ctx) error
}

package product

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateProduct(c *fiber.Ctx) error
	GetAllProduct(c *fiber.Ctx) error
}

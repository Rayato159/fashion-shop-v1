package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/product"
)

func MapProductRoute(r fiber.Router, h product.Handler) {
	r.Get("/", h.GetAllProduct)
	r.Post("/create", h.CreateProduct)
}

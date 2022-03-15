package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/color"
)

func MapColorRoute(r fiber.Router, h color.Handler) {
	r.Get("/", h.GetAllColor)
	r.Get("/:color", h.GetColorByKey)
}

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/pattern"
)

func MapPatternRoute(r fiber.Router, h pattern.Handler) {
	r.Get("/", h.GetAllPattern)
}

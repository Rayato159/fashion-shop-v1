package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/pattern"
)

func MapPatternRoute(r fiber.Router, h pattern.Handler) {
	r.Get("/", h.GetAllPattern)
	r.Get("/:pattern", h.GetPatternByKey)
	r.Post("/create", h.CreatePattern)
	r.Post("/create-bulk", h.CreatePatternButBulk)
	r.Delete("/:pattern/delete", h.DeletePattern)
}

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/category"
)

func MapCategoryRoute(r fiber.Router, h category.Handler) {
	r.Get("/", h.GetAllCategory)
	r.Post("/create", h.CreateCategory)
}

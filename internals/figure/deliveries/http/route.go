package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/figure"
)

func MapFigureRoute(r fiber.Router, h figure.Handler) {
	r.Get("/", h.GetAllFigure)
	r.Get("/:figure", h.GetFigureByKey)
	r.Post("/create", h.CreateFigure)
	r.Post("/create-bulk", h.CreateFigureButBulk)
	r.Delete("/:figure/delete", h.DeleteFigure)
}

package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	colorHttp "github.com/rayato159/fashion-shop/v1/internals/color/deliveries/http"
	colorRepository "github.com/rayato159/fashion-shop/v1/internals/color/repositories"
	colorUseCase "github.com/rayato159/fashion-shop/v1/internals/color/usecases"
	"github.com/rayato159/fashion-shop/v1/internals/models"
)

func (s *Server) MapHandlers(a *fiber.App) error {
	a.Use(logger.New())
	api := a.Group("/api")

	colorGroup := api.Group("/colors")
	colorRepo := colorRepository.NewColorRepository(s.db)
	colorUC := colorUseCase.NewColorUseCase(colorRepo)
	colorHandler := colorHttp.NewColorHandler(colorUC)
	colorHttp.MapColorRoute(colorGroup, colorHandler)

	a.Use(func(c *fiber.Ctx) error {
		log.Println("error, endpoint not found.")
		return c.Status(404).JSON(&models.Error{
			Code:    404,
			Message: "Not Found",
		})
	})

	return nil
}

package server

import (
	"log"

	"github.com/rayato159/fashion-shop/v1/internals/models"

	figureHttp "github.com/rayato159/fashion-shop/v1/internals/figure/deliveries/http"
	figureRepository "github.com/rayato159/fashion-shop/v1/internals/figure/repositories"
	figureUseCase "github.com/rayato159/fashion-shop/v1/internals/figure/usecases"

	colorHttp "github.com/rayato159/fashion-shop/v1/internals/color/deliveries/http"
	colorRepository "github.com/rayato159/fashion-shop/v1/internals/color/repositories"
	colorUseCase "github.com/rayato159/fashion-shop/v1/internals/color/usecases"

	patternHttp "github.com/rayato159/fashion-shop/v1/internals/pattern/deliveries/http"
	patternRepository "github.com/rayato159/fashion-shop/v1/internals/pattern/repositories"
	patternUseCase "github.com/rayato159/fashion-shop/v1/internals/pattern/usecases"

	categoryHttp "github.com/rayato159/fashion-shop/v1/internals/category/deliveries/http"
	categoryRepository "github.com/rayato159/fashion-shop/v1/internals/category/repositories"
	categoryUseCase "github.com/rayato159/fashion-shop/v1/internals/category/usecases"

	productHttp "github.com/rayato159/fashion-shop/v1/internals/product/deliveries/http"
	productRepository "github.com/rayato159/fashion-shop/v1/internals/product/repositories"
	productUseCase "github.com/rayato159/fashion-shop/v1/internals/product/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *Server) MapHandlers(a *fiber.App) error {
	a.Use(logger.New())
	api := a.Group("/api")

	colorGroup := api.Group("/colors")
	colorRepo := colorRepository.NewColorRepository(s.db)
	colorUC := colorUseCase.NewColorUseCase(colorRepo)
	colorHandler := colorHttp.NewColorHandler(colorUC)
	colorHttp.MapColorRoute(colorGroup, colorHandler)

	patternGroup := api.Group("/patterns")
	patternRepo := patternRepository.NewPatternRepository(s.db)
	patternUC := patternUseCase.NewPatternUseCase(patternRepo)
	patternHandler := patternHttp.NewPatternHandler(patternUC)
	patternHttp.MapPatternRoute(patternGroup, patternHandler)

	figureGroup := api.Group("/figures")
	figureRepo := figureRepository.NewFigureRepository(s.db)
	figureUC := figureUseCase.NewFigureUseCase(figureRepo)
	figureHandler := figureHttp.NewFigureHandler(figureUC)
	figureHttp.MapFigureRoute(figureGroup, figureHandler)

	categoryGroup := api.Group("/categories")
	categoryRepo := categoryRepository.NewCategoryRepository(s.db)
	categoryUC := categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler := categoryHttp.NewCategoryHandler(categoryUC)
	categoryHttp.MapCategoryRoute(categoryGroup, categoryHandler)

	productGroup := api.Group("/products")
	productRepo := productRepository.NewProductRepository(s.db)
	productUC := productUseCase.NewProductUseCase(productRepo)
	productHandler := productHttp.NewProductHandler(productUC)
	productHttp.MapProductRoute(productGroup, productHandler)

	a.Use(func(c *fiber.Ctx) error {
		log.Println("error, endpoint not found.")
		return c.Status(404).JSON(&models.Error{
			Code:    404,
			Message: "Not Found",
		})
	})

	return nil
}

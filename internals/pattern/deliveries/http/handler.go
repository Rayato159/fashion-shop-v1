package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayato159/fashion-shop/v1/internals/pattern"
)

type patternHandler struct {
	patternUC pattern.UseCase
}

func NewPatternHandler(patternUC pattern.UseCase) pattern.Handler {
	return &patternHandler{patternUC: patternUC}
}

func (h *patternHandler) GetAllPattern(c *fiber.Ctx) error {
	return nil
}

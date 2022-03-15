package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	fiber *fiber.App
	db    *sqlx.DB
}

func NewServer(db *sqlx.DB) *Server {
	return &Server{
		fiber: fiber.New(),
		db:    db,
	}
}

func (s *Server) Start() error {
	if err := s.MapHandlers(s.fiber); err != nil {
		return err
	}
	s.fiber.Listen(":3000")
	return nil
}

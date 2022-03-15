package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rayato159/fashion-shop/v1/internals/pattern"
)

type patternRepo struct {
	db *sqlx.DB
}

func NewPatternRepository(db *sqlx.DB) pattern.Repository {
	return &patternRepo{db: db}
}

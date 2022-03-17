package models

import "time"

type Figure struct {
	FigureId  int       `db:"figure_id" json:"figure_id"`
	Figure    string    `db:"figure" json:"figure"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateFigure struct {
	Figure string `json:"figure" form:"figure"`
}

type FigureFilter struct {
	Figure string `json:"figure" query:"figure"`
}

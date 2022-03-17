package models

import "time"

type Color struct {
	ColorId   int       `db:"color_id" json:"color_id"`
	Color     string    `db:"color" json:"color"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateColor struct {
	Color string `json:"color" form:"color"`
}

type ColorFilter struct {
	Color string `json:"color" query:"color"`
}

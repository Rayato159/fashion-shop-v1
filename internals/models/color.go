package models

import "time"

type Color struct {
	ColorId   int       `db:"color_id" json:"color_id"`
	Color     string    `db:"color" json:"color"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type FindOneColor struct {
	Color string `db:"color" json:"color"`
}

type ColorFilter struct {
	Color string `db:"color" json:"color"`
}
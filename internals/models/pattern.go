package models

import "time"

type Pattern struct {
	PatternId int       `db:"pattern_id" json:"pattern_id"`
	Pattern   string    `db:"pattern" json:"pattern"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreatePattern struct {
	Pattern string `db:"pattern" json:"pattern"`
}

type FindOnePattern struct {
	Pattern string `db:"pattern" json:"pattern"`
}

type FilterPattern struct {
	Pattern string `db:"pattern" json:"pattern"`
}

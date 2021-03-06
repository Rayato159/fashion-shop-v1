package models

type Category struct {
	CategoryId int        `db:"category_id" json:"category_id"`
	Size       SizeEnum   `db:"size" json:"size"`
	Gender     GenderEnum `db:"gender" json:"gender"`
	Price      float32    `db:"price" json:"price"`
	ColorId    int        `db:"color_id" json:"color_id"`
	PatternId  int        `db:"pattern_id" json:"pattern_id"`
	FigureId   int        `db:"figure_id" json:"figure_id"`
}

type CreateCategory struct {
	Gender  GenderEnum `json:"gender" form:"gender"`
	Size    SizeEnum   `json:"size" form:"size"`
	Price   float32    `json:"price" form:"price"`
	Color   string     `json:"color" form:"color"`
	Pattern string     `json:"pattern" form:"pattern"`
	Figure  string     `json:"figure" form:"figure"`
}

// enum
type SizeEnum string

const (
	XS SizeEnum = "xs"
	S  SizeEnum = "s"
	M  SizeEnum = "m"
	L  SizeEnum = "l"
	XL SizeEnum = "xl"
)

type GenderEnum string

const (
	Male   GenderEnum = "male"
	Female GenderEnum = "female"
)

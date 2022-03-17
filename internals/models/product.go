package models

import "time"

type Product struct {
	ProductId  int        `db:"product_id" json:"product_id"`
	CategoryId int        `db:"category_id" json:"category_id"`
	Gender     GenderEnum `db:"gender" json:"gender"`
	Size       SizeEnum   `db:"size" json:"size"`
	Price      float32    `db:"price" json:"price"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
}

type CreateProduct struct {
	CategoryId int `json:"category_id" form:"category_id"`
}

type ProductFilter struct {
	CategoryId int        `json:"category_id" query:"category_id"`
	Gender     GenderEnum `json:"gender" query:"gender"`
	Size       SizeEnum   `json:"size" query:"size"`
}

package model

import (
	_ "github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	Name        string         `json:"name"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Image       string         `json:"image"`
}

type CreateProductDTO struct {
	Name        string         `json:"name" validate:"required"`
	Tags        pq.StringArray `json:"tags"`
	Description string         `json:"description" validate:"required"`
	Price       float64        `json:"price" validate:"required,numeric,gte=0"`
}

type UpdateProductDTO struct {
	Name        string         `json:"name"`
	Tags        pq.StringArray `json:"tags"`
	Description string         `json:"description"`
	Price       float64        `json:"price" validate:"numeric,gte=0"`
}

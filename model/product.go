package model

import (
	"time"

	_ "github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Image       string         `json:"image"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
}

type CreateProductDTO struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Price       float64        `json:"price" validate:"required,numeric,gte=0"`
	Tags        pq.StringArray `json:"tags"`
}

type UpdateProductDTO struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price" validate:"numeric,gte=0"`
	Tags        pq.StringArray `json:"tags"`
}

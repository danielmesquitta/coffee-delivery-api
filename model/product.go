package model

import (
	"time"

	_ "github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Price       float64        `json:"price,omitempty"`
	Image       string         `json:"image,omitempty"`
	Tags        pq.StringArray `json:"tags,omitempty" gorm:"type:text[]"`
}

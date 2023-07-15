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
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Image       string         `json:"image"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
}

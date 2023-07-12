package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	Name        string         `json:"name"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
	Description string         `json:"description"`
	Price       uint           `json:"price"`
	Image       string         `json:"image"`
}

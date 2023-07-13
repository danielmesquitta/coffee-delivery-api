package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Products  []Product      `json:"products" gorm:"many2many:order_products;"`
	UserID    uint           `json:"userId"`
	User      User           `json:"user"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	Products  []*Product     `json:"products,omitempty" gorm:"many2many:order_products;"`
	UserID    uint           `json:"userId,omitempty"`
	User      *User          `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  float64        `json:"password"`
	Avatar    string         `json:"avatar"`
	AddressID uint           `json:"addressId"`
	Address   Address        `json:"address"`
}

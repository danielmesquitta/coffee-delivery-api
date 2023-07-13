package model

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	ZipCode      string         `json:"zipCode"`
	Street       string         `json:"street"`
	Number       string         `json:"number"`
	Complement   string         `json:"complement"`
	Neighborhood string         `json:"neighborhood"`
	City         string         `json:"city"`
	State        string         `json:"state"`
}

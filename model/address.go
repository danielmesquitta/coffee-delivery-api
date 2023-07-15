package model

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID           uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	UpdatedAt    time.Time      `json:"updatedAt,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	ZipCode      string         `json:"zipCode,omitempty"`
	Street       string         `json:"street,omitempty"`
	Number       string         `json:"number,omitempty"`
	Complement   string         `json:"complement,omitempty"`
	Neighborhood string         `json:"neighborhood,omitempty"`
	City         string         `json:"city,omitempty"`
	State        string         `json:"state,omitempty"`
}

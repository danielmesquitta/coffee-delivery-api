package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	CreditCard PaymentMethod = "DEBIT_CARD"
	DebitCard  PaymentMethod = "CREDIT_CARD"
	Cash       PaymentMethod = "CASH"
)

type User struct {
	ID            uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"createdAt,omitempty"`
	UpdatedAt     time.Time      `json:"updatedAt,omitempty"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	PaymentMethod PaymentMethod  `json:"paymentMethod,omitempty" sql:"type:ENUM('DEBIT_CARD', 'CREDIT_CARD', 'CASH')"`
	AddressID     uint           `json:"addressId,omitempty"`
	Address       *Address       `json:"address,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

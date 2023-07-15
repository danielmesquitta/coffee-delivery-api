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
	ID            uint           `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	PaymentMethod PaymentMethod  `json:"paymentMethod" sql:"type:ENUM('DEBIT_CARD', 'CREDIT_CARD', 'CASH')"`
	AddressID     uint           `json:"addressId"`
	Address       Address        `json:"address" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

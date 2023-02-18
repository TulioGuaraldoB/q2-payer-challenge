package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint `gorm:"primaryKey"`
	PayerID     uint
	PayerWallet Wallet `gorm:"foreignKey:PayerID"`
	Amount      float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

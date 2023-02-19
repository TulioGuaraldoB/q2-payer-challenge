package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID               uint `gorm:"primaryKey"`
	PayerWalletID    uint
	PayerWallet      Wallet `gorm:"foreignKey:PayerWalletID"`
	ReceiverWalletID uint
	ReceiverWallet   Wallet `gorm:"foreignKey:ReceiverWalletID"`
	Amount           float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

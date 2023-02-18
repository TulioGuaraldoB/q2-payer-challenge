package migration

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Wallet{})
	db.AutoMigrate(&model.Transaction{})
}

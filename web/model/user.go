package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FullName  string `gorm:"size:255"`
	CPF       string `gorm:"size:11;unique"`
	CNPJ      string `gorm:"size:14"`
	Email     string `gorm:"size:175;unique"`
	Password  string `gorm:"size:100"`
	UserType  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

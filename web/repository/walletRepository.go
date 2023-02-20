package repository

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"gorm.io/gorm"
)

type IWalletRepository interface {
	GetWalletByUserId(userId uint) (*model.Wallet, error)
	CreateWallet(wallet *model.Wallet) error
	DeleteWallet(walletId uint) error
	UpdateWalletBalance(walletId uint, newBalance float64) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) IWalletRepository {
	return &walletRepository{
		db: db,
	}
}

func (r *walletRepository) GetWalletByUserId(userId uint) (*model.Wallet, error) {
	wallet := new(model.Wallet)
	if err := r.db.Joins("User").
		Where("user_id = ?", userId).
		First(&wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *walletRepository) CreateWallet(wallet *model.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *walletRepository) DeleteWallet(walletId uint) error {
	wallet := new(model.Wallet)
	return r.db.Where("id = ?", walletId).Delete(&wallet).Error
}

func (r *walletRepository) UpdateWalletBalance(walletId uint, newBalance float64) error {
	wallet := new(model.Wallet)

	tx := r.db.Begin()

	if err := tx.Model(wallet).
		Where("id = ?", &walletId).
		Update("balance", newBalance).
		Error; err != nil {
		r.db.Rollback()
		return err
	}

	return tx.Commit().Error
}

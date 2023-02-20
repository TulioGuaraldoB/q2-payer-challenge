package repository

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	GetTransactionById(transactionId uint) (*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) GetTransactionById(transactionId uint) (*model.Transaction, error) {
	transaction := new(model.Transaction)
	if err := r.db.
		Preload("PayerWallet.User").
		Preload("ReceiverWallet.User").
		First(&transaction, &transactionId).
		Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

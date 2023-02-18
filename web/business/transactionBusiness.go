package business

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository"
)

type ITransactionBusiness interface {
	GetTransactionById(transactionId uint) (*model.Transaction, error)
}

type transactionBusiness struct {
	transactionRepository repository.ITransactionRepository
}

func NewTransactionBusiness(transactionRepository repository.ITransactionRepository) ITransactionBusiness {
	return &transactionBusiness{
		transactionRepository: transactionRepository,
	}
}

func (b *transactionBusiness) GetTransactionById(transactionId uint) (*model.Transaction, error) {
	return b.transactionRepository.GetTransactionById(transactionId)
}

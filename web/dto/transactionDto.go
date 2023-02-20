package dto

import "github.com/TulioGuaraldoB/q2-payer-challenge/web/model"

type TransactionResponse struct {
	ID               uint           `json:"id"`
	PayerWalletID    uint           `json:"payer_id"`
	PayerWallet      WalletResponse `json:"payer_wallet"`
	ReceiverWalletID uint           `json:"receiver_wallet_id"`
	ReceiverWallet   WalletResponse `json:"receiver_wallet"`
	Amount           float64        `json:"amount"`
}

type TransactionRequest struct {
	PayerWalletUserID    uint    `json:"payer_id"`
	ReceiverWalletUserID uint    `json:"receiver_id"`
	Amount               float64 `json:"amount"`
}

func ParseTransactionToResponse(transaction *model.Transaction) *TransactionResponse {
	return &TransactionResponse{
		ID:            transaction.ID,
		PayerWalletID: transaction.PayerWalletID,
		PayerWallet: WalletResponse{
			ID:     transaction.PayerWallet.ID,
			UserID: transaction.PayerWallet.UserID,
			User: UserResponse{
				ID:       transaction.PayerWallet.User.ID,
				FullName: transaction.PayerWallet.User.FullName,
				CPF:      transaction.PayerWallet.User.CPF,
				CNPJ:     transaction.PayerWallet.User.CNPJ,
				Email:    transaction.PayerWallet.User.Email,
				UserType: transaction.PayerWallet.User.UserType,
			},
			Balance: transaction.PayerWallet.Balance,
		},
		ReceiverWalletID: transaction.ReceiverWalletID,
		ReceiverWallet: WalletResponse{
			ID:     transaction.ReceiverWallet.ID,
			UserID: transaction.ReceiverWallet.UserID,
			User: UserResponse{
				ID:       transaction.ReceiverWallet.User.ID,
				FullName: transaction.ReceiverWallet.User.FullName,
				CPF:      transaction.ReceiverWallet.User.CPF,
				CNPJ:     transaction.ReceiverWallet.User.CNPJ,
				Email:    transaction.ReceiverWallet.User.Email,
				UserType: transaction.ReceiverWallet.User.UserType,
			},
			Balance: transaction.ReceiverWallet.Balance,
		},
		Amount: transaction.Amount,
	}
}

func ParseRequestToTransaction(transactionRequest *TransactionRequest) *model.Transaction {
	return &model.Transaction{
		PayerWalletID:    transactionRequest.PayerWalletUserID,
		ReceiverWalletID: transactionRequest.ReceiverWalletUserID,
		Amount:           transactionRequest.Amount,
	}
}

package dto

import "github.com/TulioGuaraldoB/q2-payer-challenge/web/model"

type TransactionResponse struct {
	ID             uint           `json:"id"`
	PayerWalletID  uint           `json:"payer_id"`
	PayerWallet    WalletResponse `json:"payer_wallet"`
	TargetWalletID uint           `json:"target_wallet_id"`
	TargetWallet   WalletResponse `json:"target_wallet"`
	Amount         float64        `json:"amount"`
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
		TargetWalletID: transaction.TargetWalletID,
		TargetWallet: WalletResponse{
			ID:     transaction.TargetWallet.ID,
			UserID: transaction.TargetWallet.UserID,
			User: UserResponse{
				ID:       transaction.TargetWallet.User.ID,
				FullName: transaction.TargetWallet.User.FullName,
				CPF:      transaction.TargetWallet.User.CPF,
				CNPJ:     transaction.TargetWallet.User.CNPJ,
				Email:    transaction.TargetWallet.User.Email,
				UserType: transaction.TargetWallet.User.UserType,
			},
			Balance: transaction.TargetWallet.Balance,
		},
		Amount: transaction.Amount,
	}
}

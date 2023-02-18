package dto

import "github.com/TulioGuaraldoB/q2-payer-challenge/web/model"

type WalletResponse struct {
	ID      uint         `json:"id"`
	UserID  uint         `json:"user_id"`
	User    UserResponse `json:"user"`
	Balance float64      `json:"balance"`
}

type WalletRequest struct {
	UserID  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}

func ParseWalletToResponse(wallet *model.Wallet) *WalletResponse {
	return &WalletResponse{
		ID:     wallet.ID,
		UserID: wallet.UserID,
		User: UserResponse{
			ID:       wallet.User.ID,
			FullName: wallet.User.FullName,
			CPF:      wallet.User.CPF,
			CNPJ:     wallet.User.CNPJ,
			Email:    wallet.User.Email,
			UserType: wallet.User.UserType,
		},
		Balance: wallet.Balance,
	}
}

func ParseRequestToWallet(walletRequest *WalletRequest) *model.Wallet {
	return &model.Wallet{
		UserID:  walletRequest.UserID,
		Balance: walletRequest.Balance,
	}
}

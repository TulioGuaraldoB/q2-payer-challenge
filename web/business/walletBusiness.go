package business

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/util/encrypt"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository"
)

type IWalletBusiness interface {
	GetWalletByUserId(userId uint) (*model.Wallet, error)
	GetWalletByUserCredentials(userCredentials *dto.UserCredentials) (*model.Wallet, error)
	CreateWallet(wallet *model.Wallet) error
	DeleteWallet(walletId uint) error
	UpdateWalletBalance(walletId uint, newBalance float64) error
}

type walletBusiness struct {
	walletRepository repository.IWalletRepository
	userRepository   repository.IUserRepository
}

func NewWalletBusiness(walletRepository repository.IWalletRepository,
	userRepository repository.IUserRepository) IWalletBusiness {
	return &walletBusiness{
		walletRepository: walletRepository,
		userRepository:   userRepository,
	}
}

func (b *walletBusiness) GetWalletByUserId(userId uint) (*model.Wallet, error) {
	return b.walletRepository.GetWalletByUserId(userId)
}

func (b *walletBusiness) GetWalletByUserCredentials(userCredentials *dto.UserCredentials) (*model.Wallet, error) {
	hashedPassword := encrypt.HashToSHA256(userCredentials.Password)
	userCredentials.Password = hashedPassword

	user, err := b.userRepository.GetUserByCredentials(userCredentials.Email, userCredentials.Password)
	if err != nil {
		return nil, err
	}

	return b.walletRepository.GetWalletByUserId(user.ID)
}

func (b *walletBusiness) CreateWallet(wallet *model.Wallet) error {
	return b.walletRepository.CreateWallet(wallet)
}

func (b *walletBusiness) DeleteWallet(walletId uint) error {
	return b.walletRepository.DeleteWallet(walletId)
}

func (b *walletBusiness) UpdateWalletBalance(walletId uint, newBalance float64) error {
	return b.walletRepository.UpdateWalletBalance(walletId, newBalance)
}

package business

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/util/encrypt"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/constant"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/service"
)

type IWalletBusiness interface {
	GetWalletByUserId(userId uint) (*model.Wallet, error)
	GetWalletByUserCredentials(userCredentials *dto.UserCredentials) (*model.Wallet, error)
	CreateWallet(wallet *model.Wallet) error
	DeleteWallet(walletId uint) error
	UpdateWalletBalance(walletId uint, newBalance float64) error
	DepositToWalletBalance(userId uint, newBalance float64) error
}

type walletBusiness struct {
	walletRepository     repository.IWalletRepository
	userRepository       repository.IUserRepository
	authorizationService service.IAuthorizationService
}

func NewWalletBusiness(walletRepository repository.IWalletRepository,
	userRepository repository.IUserRepository,
	authorizationService service.IAuthorizationService) IWalletBusiness {
	return &walletBusiness{
		walletRepository:     walletRepository,
		userRepository:       userRepository,
		authorizationService: authorizationService,
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

func (b *walletBusiness) DepositToWalletBalance(userId uint, newBalance float64) error {
	wallet, err := b.walletRepository.GetWalletByUserId(userId)
	if err != nil {
		return err
	}

	oldBalance := wallet.Balance
	newWalletBalance := oldBalance + newBalance
	return b.UpdateWalletBalance(userId, newWalletBalance)
}

func (b *walletBusiness) PaymentWalletTransaction(transactionRequest *dto.TransactionRequest, payerWalletUserId, targetWalletUserId uint, paymentTransactionValue float64) error {
	wallet, err := b.walletRepository.GetWalletByUserId(transactionRequest.PayerWalletUserID)
	if err != nil {
		return err
	}

	if wallet.Balance < paymentTransactionValue {
		return constant.INSUFFICIENT_FUND
	}

	authorizerResponse, err := b.authorizationService.CheckAuthorizerApi()
	if err != nil {
		authorizerResponse.Authorized = false
	}

	if !authorizerResponse.Authorized {
		return constant.UNAUTHORIZED_TRANSACTION
	}

	newPayerBalance := (wallet.Balance - paymentTransactionValue)
	if err := b.walletRepository.UpdateWalletBalance(wallet.ID, newPayerBalance); err != nil {
		return err
	}

	targetWallet, err := b.walletRepository.GetWalletByUserId(transactionRequest.ReceiverWalletUsertID)
	if err != nil {
		return err
	}

	newTargetBalance := (targetWallet.Balance + paymentTransactionValue)
	return b.walletRepository.UpdateWalletBalance(targetWallet.ID, newTargetBalance)
}

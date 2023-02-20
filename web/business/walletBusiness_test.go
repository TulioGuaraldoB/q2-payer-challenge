package business_test

import (
	"net/http"
	"testing"

	"github.com/TulioGuaraldoB/q2-payer-challenge/util/encrypt"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/business"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository/mock"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/service"
	faker "github.com/brianvoe/gofakeit"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type walletBusinessTest struct {
	description    string
	setWalletMocks func(*mock.MockIWalletRepository)
	setUserMocks   func(*mock.MockIUserRepository)
	isErrExpected  bool
}

// GetWalletByUserId
func TestGetWalletByUserIdBusiness(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	tests := []walletBusinessTest{
		{
			description: "Should return no error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					GetWalletByUserId(mockWallet.UserID).
					Return(&mockWallet, nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					GetWalletByUserId(mockWallet.UserID).
					Return(nil, gorm.ErrRecordNotFound)
			},
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			miwr := mock.NewMockIWalletRepository(controller)
			miur := mock.NewMockIUserRepository(controller)
			mias := service.NewAuthorizationService(&http.Client{})

			testCase.setWalletMocks(miwr)

			// Act
			walletBusiness := business.NewWalletBusiness(miwr, miur, mias)
			_, err := walletBusiness.GetWalletByUserId(mockWallet.UserID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetWalletByUserCredentials
func TestGetWalletByUserCredentialsBusiness(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	mockUserCredentials := dto.UserCredentials{}
	faker.Struct(&mockUserCredentials)

	mockUser := model.User{}
	faker.Struct(&mockUser)

	mockUser.ID = 10
	mockWallet.UserID = mockUser.ID

	tests := []walletBusinessTest{
		{
			description: "Should return no error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					GetWalletByUserId(mockWallet.UserID).
					Return(&mockWallet, nil)
			},
			setUserMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByCredentials(mockUserCredentials.Email, encrypt.HashToSHA256(mockUserCredentials.Password)).
					Return(&mockUser, nil)

			},
			isErrExpected: false,
		},
		{
			description:    "Should return error on get user by credentials",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {},
			setUserMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByCredentials(mockUserCredentials.Email, encrypt.HashToSHA256(mockUserCredentials.Password)).
					Return(nil, gorm.ErrRecordNotFound)

			},
			isErrExpected: true,
		},
		{
			description: "Should return error on get wallet",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					GetWalletByUserId(mockWallet.UserID).
					Return(nil, gorm.ErrRecordNotFound)
			},
			setUserMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByCredentials(mockUserCredentials.Email, encrypt.HashToSHA256(mockUserCredentials.Password)).
					Return(&mockUser, nil)

			},
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			miwr := mock.NewMockIWalletRepository(controller)
			miur := mock.NewMockIUserRepository(controller)
			mias := service.NewAuthorizationService(&http.Client{})

			testCase.setWalletMocks(miwr)
			testCase.setUserMocks(miur)

			// Act
			walletBusiness := business.NewWalletBusiness(miwr, miur, mias)
			_, err := walletBusiness.GetWalletByUserCredentials(&mockUserCredentials)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// CreateWallet
func TestCreateWalletBusiness(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	tests := []walletBusinessTest{
		{
			description: "Should return no error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					CreateWallet(&mockWallet).
					Return(nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					CreateWallet(&mockWallet).
					Return(gorm.ErrInvalidValue)
			},
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			miwr := mock.NewMockIWalletRepository(controller)
			miur := mock.NewMockIUserRepository(controller)
			mias := service.NewAuthorizationService(&http.Client{})

			testCase.setWalletMocks(miwr)

			// Act
			walletBusiness := business.NewWalletBusiness(miwr, miur, mias)
			err := walletBusiness.CreateWallet(&mockWallet)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// DeleteWallet
func TestDeleteWalletBusiness(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	tests := []walletBusinessTest{
		{
			description: "Should return no error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					DeleteWallet(mockWallet.ID).
					Return(nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					DeleteWallet(mockWallet.ID).
					Return(gorm.ErrInvalidValue)
			},
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			miwr := mock.NewMockIWalletRepository(controller)
			miur := mock.NewMockIUserRepository(controller)
			mias := service.NewAuthorizationService(&http.Client{})

			testCase.setWalletMocks(miwr)

			// Act
			walletBusiness := business.NewWalletBusiness(miwr, miur, mias)
			err := walletBusiness.DeleteWallet(mockWallet.ID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// UpdateWalletBalance
func TestUpdateWalletBalanceBusiness(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	tests := []walletBusinessTest{
		{
			description: "Should return no error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					UpdateWalletBalance(mockWallet.ID, mockWallet.Balance).
					Return(nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					UpdateWalletBalance(mockWallet.ID, mockWallet.Balance).
					Return(gorm.ErrInvalidValue)
			},
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			miwr := mock.NewMockIWalletRepository(controller)
			miur := mock.NewMockIUserRepository(controller)
			mias := service.NewAuthorizationService(&http.Client{})

			testCase.setWalletMocks(miwr)

			// Act
			walletBusiness := business.NewWalletBusiness(miwr, miur, mias)
			err := walletBusiness.UpdateWalletBalance(mockWallet.ID, mockWallet.Balance)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// DepositToWalletBalance
func TestDepositToWalletBalanceBusiness(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	mockValue := faker.Float64()

	tests := []walletBusinessTest{
		{
			description: "Should return no error",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					GetWalletByUserId(mockWallet.UserID).
					Return(&mockWallet, nil)
				mir.EXPECT().
					UpdateWalletBalance(mockWallet.ID, (mockWallet.Balance + mockValue)).
					Return(nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error wallet not found",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					GetWalletByUserId(mockWallet.UserID).
					Return(nil, gorm.ErrRecordNotFound)
			},
			isErrExpected: true,
		},
		{
			description: "Should return error on update wallet value",
			setWalletMocks: func(mir *mock.MockIWalletRepository) {
				mir.EXPECT().
					GetWalletByUserId(mockWallet.UserID).
					Return(&mockWallet, nil)
				mir.EXPECT().
					UpdateWalletBalance(mockWallet.ID, (mockWallet.Balance + mockValue)).
					Return(gorm.ErrInvalidValue)
			},
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			miwr := mock.NewMockIWalletRepository(controller)
			miur := mock.NewMockIUserRepository(controller)
			mias := service.NewAuthorizationService(&http.Client{})

			testCase.setWalletMocks(miwr)

			// Act
			walletBusiness := business.NewWalletBusiness(miwr, miur, mias)
			err := walletBusiness.DepositToWalletBalance(mockWallet.ID, mockValue)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// PaymentWalletTransaction
func TestPaymentWalletTransactionBusiness(t *testing.T) {

}

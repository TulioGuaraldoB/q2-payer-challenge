package business_test

import (
	"testing"

	"github.com/TulioGuaraldoB/q2-payer-challenge/web/business"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository/mock"
	faker "github.com/brianvoe/gofakeit"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type transactionBusinessTest struct {
	description   string
	setMocks      func(*mock.MockITransactionRepository)
	isErrExpected bool
}

// GetTransactionById
func TestGetTransactionByIdBusiness(t *testing.T) {
	mockTransaction := model.Transaction{}
	faker.Struct(&mockTransaction)

	tests := []transactionBusinessTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockITransactionRepository) {
				mir.EXPECT().
					GetTransactionById(mockTransaction.ID).
					Return(&mockTransaction, nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setMocks: func(mir *mock.MockITransactionRepository) {
				mir.EXPECT().
					GetTransactionById(mockTransaction.ID).
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

			mitr := mock.NewMockITransactionRepository(controller)
			testCase.setMocks(mitr)

			// Act
			transactionBusiness := business.NewTransactionBusiness(mitr)
			_, err := transactionBusiness.GetTransactionById(mockTransaction.ID)
			if !testCase.isErrExpected {
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

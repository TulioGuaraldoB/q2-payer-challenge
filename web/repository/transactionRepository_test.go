package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository/mock"
	faker "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

type transactionRepositoryTest struct {
	description       string
	expectedQuery     string
	expectedPostQuery string
	isErrExpected     bool
}

// GetTransactionById
func TestGetTransactionByIdRepository(t *testing.T) {
	mockTransaction := model.Transaction{}
	faker.Struct(&mockTransaction)

	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	mockTransaction.PayerWalletID = mockWallet.ID
	mockTransaction.ReceiverWalletID = mockWallet.ID + 10

	mockUser := model.User{}
	faker.Struct(&mockUser)

	mockUser.ID = mockWallet.UserID

	tests := []transactionRepositoryTest{
		{
			description:       "Should return no error",
			expectedQuery:     "SELECT * FROM `transactions` WHERE `transactions`.`id` = ? AND `transactions`.`deleted_at` IS NULL ORDER BY `transactions`.`id` LIMIT 1",
			expectedPostQuery: "SELECT * FROM `wallets` WHERE `wallets`.`id` = ? AND `wallets`.`deleted_at` IS NULL",
			isErrExpected:     false,
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT `wallets`id`,`wallets``user_id`,`wallets``balance`,`wallets``created_at`,`wallets``updated_at`,`wallets``deleted_at`,`User``id` AS `User__id`,`User`.`full_name` AS `User__full_name`,`User`.`cpf` AS `User__cpf`,`User`.`cnpj` AS `User__cnpj`,`User`.`email` AS `User__email`,`User`.`password` AS `User__password`,`User`.`user_type` AS `User__user_type`,`User`.`created_at` AS `User__created_at`,`User`.`updated_at` AS `User__updated_at`,`User`.`deleted_at` AS `User__deleted_at` FROM `wallets` LEFT JOIN `users` `User` ON `wallets`.`user_id` = `User`.`id` AND `User`.`deleted_at` IS NULL WHERE user_id = ? AND `wallets`.`deleted_at` IS NULL ORDER BY `wallets`.`id` LIMIT 1",
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			rows := sqlmock.NewRows([]string{
				"ID",
				"PayerWalletID",
				"ReceiverWalletID",
				"Amount",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockTransaction.ID,
				mockTransaction.PayerWalletID,
				mockTransaction.ReceiverWalletID,
				mockTransaction.Amount,
				mockTransaction.CreatedAt,
				mockTransaction.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			postQuery := regexp.QuoteMeta(testCase.expectedPostQuery)
			sqlMock.ExpectQuery(query).WillReturnRows(rows)
			sqlMock.ExpectQuery(postQuery).WillReturnRows(rows)

			// Act
			transactionRepository := repository.NewTransactionRepository(mockDb)
			_, err := transactionRepository.GetTransactionById(mockWallet.UserID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

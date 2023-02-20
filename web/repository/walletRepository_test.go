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

type walletRepositoryTest struct {
	description       string
	expectedQuery     string
	expectedPostQuery string
	expectedRollback  bool
	isErrExpected     bool
}

// GetWalletByUserId
func TestGetWalletByUserIdRepository(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	mockUser := model.User{}
	faker.Struct(&mockUser)

	mockUser.ID = mockWallet.UserID

	tests := []walletRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT `wallets`.`id`,`wallets`.`user_id`,`wallets`.`balance`,`wallets`.`created_at`,`wallets`.`updated_at`,`wallets`.`deleted_at`,`User`.`id` AS `User__id`,`User`.`full_name` AS `User__full_name`,`User`.`cpf` AS `User__cpf`,`User`.`cnpj` AS `User__cnpj`,`User`.`email` AS `User__email`,`User`.`password` AS `User__password`,`User`.`user_type` AS `User__user_type`,`User`.`created_at` AS `User__created_at`,`User`.`updated_at` AS `User__updated_at`,`User`.`deleted_at` AS `User__deleted_at` FROM `wallets` LEFT JOIN `users` `User` ON `wallets`.`user_id` = `User`.`id` AND `User`.`deleted_at` IS NULL WHERE user_id = ? AND `wallets`.`deleted_at` IS NULL ORDER BY `wallets`.`id` LIMIT 1",
			isErrExpected: false,
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
				"UserID",
				"Balance",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockWallet.ID,
				mockWallet.UserID,
				mockWallet.Balance,
				mockWallet.CreatedAt,
				mockWallet.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectQuery(query).WillReturnRows(rows)

			// Act
			walletRepository := repository.NewWalletRepository(mockDb)
			_, err := walletRepository.GetWalletByUserId(mockWallet.UserID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// CreateWallet
func TestCreateWalletRepository(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	tests := []walletRepositoryTest{
		{
			description:       "Should return no error",
			expectedQuery:     "INSERT INTO `users` (`full_name`,`cpf`,`cnpj`,`email`,`password`,`user_type`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`",
			expectedPostQuery: "INSERT INTO `wallets` (`user_id`,`balance`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?)",
			isErrExpected:     false,
		},
		{
			description:   "Should return error",
			expectedQuery: "INSERT INTO `users` (`full_name``cpf``cnpj``email``password``user_type``created_at``updated_at``deleted_at`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATEKEYUPDATE `id=`id`",
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectBegin()
			sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))

			if len(testCase.expectedPostQuery) > 0 {
				postQuery := regexp.QuoteMeta(testCase.expectedPostQuery)
				sqlMock.ExpectExec(postQuery).WillReturnResult(sqlmock.NewResult(1, 1))
			}

			sqlMock.ExpectCommit()
			sqlMock.ExpectClose()

			// Act
			walletRepository := repository.NewWalletRepository(mockDb)
			err := walletRepository.CreateWallet(&mockWallet)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// DeleteWallet
func TestDeleteWalletRepository(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	tests := []walletRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "UPDATE `wallets` SET `deleted_at`=? WHERE id = ? AND `wallets`.`deleted_at` IS NULL",
			isErrExpected: false,
		},
		{
			description:   "Should return error",
			expectedQuery: "UPDATE `wallets` SET `deleted_at`=WHERE id=AND `wallets`.`deleted_at` IS NULL",
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectBegin()
			sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			sqlMock.ExpectCommit()
			sqlMock.ExpectClose()

			// Act
			walletRepository := repository.NewWalletRepository(mockDb)
			err := walletRepository.DeleteWallet(mockWallet.ID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// UpdateWalletBalance
func TestUpdateWalletBalanceRepository(t *testing.T) {
	mockWallet := model.Wallet{}
	faker.Struct(&mockWallet)

	tests := []walletRepositoryTest{
		{
			description:      "Should return no error",
			expectedQuery:    "UPDATE `wallets` SET `balance`=?,`updated_at`=? WHERE id = ? AND `wallets`.`deleted_at",
			expectedRollback: false,
			isErrExpected:    false,
		},
		{
			description:      "Should return no error on rollback",
			expectedQuery:    "UPDATE `wallets` SET `balance`=?,`updated_at`=? WHERE id = ? AND `wallets`.`deleted_at",
			expectedRollback: true,
			isErrExpected:    false,
		},
		{
			description:      "Should return error",
			expectedQuery:    "UPDATE `wallets` SET `balance`=,`updated_at`=WHERE id=AND `wallets`.`deleted_at",
			expectedRollback: false,
			isErrExpected:    true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectBegin()
			sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			sqlMock.ExpectCommit()
			sqlMock.ExpectClose()

			// Act
			walletRepository := repository.NewWalletRepository(mockDb)
			err := walletRepository.UpdateWalletBalance(mockWallet.ID, mockWallet.Balance)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

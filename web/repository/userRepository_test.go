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

type userRepositoryTest struct {
	description   string
	expectedQuery string
	isErrExpected bool
}

// GetAllUsers
func TestGetAllUsersRepository(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL",
			isErrExpected: false,
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `users` WHERE`users``deleted_at` IS NULL",
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			rows := sqlmock.NewRows([]string{
				"ID",
				"FullName",
				"CPF",
				"CNPJ",
				"Email",
				"Password",
				"UserType",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockUser.ID,
				mockUser.FullName,
				mockUser.CPF,
				mockUser.CNPJ,
				mockUser.Email,
				mockUser.Password,
				mockUser.UserType,
				mockUser.CreatedAt,
				mockUser.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectQuery(query).WillReturnRows(rows)

			// Act
			userRepository := repository.NewUserRepository(mockDb)
			_, err := userRepository.GetAllUsers()
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetUserById
func TestGetUserByIdRepository(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
			isErrExpected: false,
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `users` WHERE`users``id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			rows := sqlmock.NewRows([]string{
				"ID",
				"FullName",
				"CPF",
				"CNPJ",
				"Email",
				"Password",
				"UserType",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockUser.ID,
				mockUser.FullName,
				mockUser.CPF,
				mockUser.CNPJ,
				mockUser.Email,
				mockUser.Password,
				mockUser.UserType,
				mockUser.CreatedAt,
				mockUser.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectQuery(query).WillReturnRows(rows)

			// Act
			userRepository := repository.NewUserRepository(mockDb)
			_, err := userRepository.GetUserById(mockUser.ID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetUserByEmail
func TestGetUserByEmailRepository(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
			isErrExpected: false,
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `users` WHERE`users``id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			rows := sqlmock.NewRows([]string{
				"ID",
				"FullName",
				"CPF",
				"CNPJ",
				"Email",
				"Password",
				"UserType",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockUser.ID,
				mockUser.FullName,
				mockUser.CPF,
				mockUser.CNPJ,
				mockUser.Email,
				mockUser.Password,
				mockUser.UserType,
				mockUser.CreatedAt,
				mockUser.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectQuery(query).WillReturnRows(rows)

			// Act
			userRepository := repository.NewUserRepository(mockDb)
			_, err := userRepository.GetUserByEmail(mockUser.Email)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetUserByCpf
func TestGetUserByCpfRepository(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `users` WHERE cpf = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
			isErrExpected: false,
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `users` WHERE`users``id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			rows := sqlmock.NewRows([]string{
				"ID",
				"FullName",
				"CPF",
				"CNPJ",
				"Email",
				"Password",
				"UserType",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockUser.ID,
				mockUser.FullName,
				mockUser.CPF,
				mockUser.CNPJ,
				mockUser.Email,
				mockUser.Password,
				mockUser.UserType,
				mockUser.CreatedAt,
				mockUser.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectQuery(query).WillReturnRows(rows)

			// Act
			userRepository := repository.NewUserRepository(mockDb)
			_, err := userRepository.GetUserByCpf(mockUser.CPF)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetUserByCredentials
func TestGetUserByCredentialsRepository(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `users` WHERE (email = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
			isErrExpected: false,
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `users` WHERE`users``id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1",
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockDb, sqlMock := mock.MockDatabase()

			rows := sqlmock.NewRows([]string{
				"ID",
				"FullName",
				"CPF",
				"CNPJ",
				"Email",
				"Password",
				"UserType",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockUser.ID,
				mockUser.FullName,
				mockUser.CPF,
				mockUser.CNPJ,
				mockUser.Email,
				mockUser.Password,
				mockUser.UserType,
				mockUser.CreatedAt,
				mockUser.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			sqlMock.ExpectQuery(query).WillReturnRows(rows)

			// Act
			userRepository := repository.NewUserRepository(mockDb)
			_, err := userRepository.GetUserByCredentials(mockUser.Email, mockUser.Password)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// CreateUser
func TestCreateUserRepository(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "INSERT INTO `users` (`full_name`,`cpf`,`cnpj`,`email`,`password`,`user_type`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?)",
			isErrExpected: false,
		},
		{
			description:   "Should return error",
			expectedQuery: "INSERT INTO `users` (`full_name``cpf``cnpj``email``password``user_type``created_at``updated_at``deleted_at`)VALUES(?,?,?,?,?,?,?,?,?)",
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
			userRepository := repository.NewUserRepository(mockDb)
			err := userRepository.CreateUser(&mockUser)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// DeleteUser
func TestDeleteUserRepository(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "UPDATE `users` SET `deleted_at`=? WHERE id = ? AND `users`.`deleted_at` IS NULL",
			isErrExpected: false,
		},
		{
			description:   "Should return error",
			expectedQuery: "INSERT INTO `users` (`full_name``cpf``cnpj``email``password``user_type``created_at``updated_at``deleted_at`)VALUES(?,?,?,?,?,?,?,?,?)",
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
			userRepository := repository.NewUserRepository(mockDb)
			err := userRepository.DeleteUser(mockUser.ID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

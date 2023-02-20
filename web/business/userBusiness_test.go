package business_test

import (
	"testing"

	"github.com/TulioGuaraldoB/q2-payer-challenge/util/encrypt"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/business"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository/mock"
	faker "github.com/brianvoe/gofakeit"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type userBusinessTest struct {
	description   string
	setMocks      func(*mock.MockIUserRepository)
	isErrExpected bool
}

// GetAllUsers
func TestGetAllUsersBusiness(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	mockUsers := []model.User{}
	mockUsers = append(mockUsers, mockUser)

	tests := []userBusinessTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetAllUsers().
					Return(mockUsers, nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetAllUsers().
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

			miur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(miur)

			// Act
			userBusiness := business.NewUserBusiness(miur)
			_, err := userBusiness.GetAllUsers()
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetUserById
func TestGetUserByIdBusiness(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userBusinessTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserById(mockUser.ID).
					Return(&mockUser, nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserById(mockUser.ID).
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

			miur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(miur)

			// Act
			userBusiness := business.NewUserBusiness(miur)
			_, err := userBusiness.GetUserById(mockUser.ID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetUserByEmail
func TestGetUserByEmailBusiness(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userBusinessTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByEmail(mockUser.Email).
					Return(&mockUser, nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByEmail(mockUser.Email).
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

			miur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(miur)

			// Act
			userBusiness := business.NewUserBusiness(miur)
			_, err := userBusiness.GetUserByEmail(mockUser.Email)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetUserByCpf
func TestGetUserByCpfBusiness(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userBusinessTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByCpf(mockUser.CPF).
					Return(&mockUser, nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByCpf(mockUser.CPF).
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

			miur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(miur)

			// Act
			userBusiness := business.NewUserBusiness(miur)
			_, err := userBusiness.GetUserByCpf(mockUser.CPF)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// GetUserByCredentials
func TestGetUserByCredentialsBusiness(t *testing.T) {
	mockUserCredentials := dto.UserCredentials{}
	faker.Struct(&mockUserCredentials)

	mockUser := model.User{}
	mockUser.Email = mockUserCredentials.Email

	tests := []userBusinessTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByCredentials(mockUserCredentials.Email, encrypt.HashToSHA256(mockUserCredentials.Password)).
					Return(&mockUser, nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByCredentials(mockUserCredentials.Email, encrypt.HashToSHA256(mockUserCredentials.Password)).
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

			miur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(miur)

			// Act
			userBusiness := business.NewUserBusiness(miur)
			_, err := userBusiness.GetUserByCredentials(&mockUserCredentials)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// CreateUser
func TestCreateUserBusiness(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userBusinessTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					CreateUser(&mockUser).
					Return(nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					CreateUser(&mockUser).
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

			miur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(miur)

			// Act
			userBusiness := business.NewUserBusiness(miur)
			err := userBusiness.CreateUser(&mockUser)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

// DeleteUser
func TestDeleteUserBusiness(t *testing.T) {
	mockUser := model.User{}
	faker.Struct(&mockUser)

	tests := []userBusinessTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					DeleteUser(mockUser.ID).
					Return(nil)
			},
			isErrExpected: false,
		},
		{
			description: "Should return error",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					DeleteUser(mockUser.ID).
					Return(gorm.ErrRecordNotFound)
			},
			isErrExpected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			miur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(miur)

			// Act
			userBusiness := business.NewUserBusiness(miur)
			err := userBusiness.DeleteUser(mockUser.ID)
			if !testCase.isErrExpected { // Assert
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}

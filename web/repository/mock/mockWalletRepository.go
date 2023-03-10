// Code generated by MockGen. DO NOT EDIT.
// Source: web/repository/walletRepository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIWalletRepository is a mock of IWalletRepository interface.
type MockIWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIWalletRepositoryMockRecorder
}

// MockIWalletRepositoryMockRecorder is the mock recorder for MockIWalletRepository.
type MockIWalletRepositoryMockRecorder struct {
	mock *MockIWalletRepository
}

// NewMockIWalletRepository creates a new mock instance.
func NewMockIWalletRepository(ctrl *gomock.Controller) *MockIWalletRepository {
	mock := &MockIWalletRepository{ctrl: ctrl}
	mock.recorder = &MockIWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWalletRepository) EXPECT() *MockIWalletRepositoryMockRecorder {
	return m.recorder
}

// CreateWallet mocks base method.
func (m *MockIWalletRepository) CreateWallet(wallet *model.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", wallet)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockIWalletRepositoryMockRecorder) CreateWallet(wallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockIWalletRepository)(nil).CreateWallet), wallet)
}

// DeleteWallet mocks base method.
func (m *MockIWalletRepository) DeleteWallet(walletId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWallet", walletId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWallet indicates an expected call of DeleteWallet.
func (mr *MockIWalletRepositoryMockRecorder) DeleteWallet(walletId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWallet", reflect.TypeOf((*MockIWalletRepository)(nil).DeleteWallet), walletId)
}

// GetWalletByUserId mocks base method.
func (m *MockIWalletRepository) GetWalletByUserId(userId uint) (*model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletByUserId", userId)
	ret0, _ := ret[0].(*model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletByUserId indicates an expected call of GetWalletByUserId.
func (mr *MockIWalletRepositoryMockRecorder) GetWalletByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletByUserId", reflect.TypeOf((*MockIWalletRepository)(nil).GetWalletByUserId), userId)
}

// UpdateWalletBalance mocks base method.
func (m *MockIWalletRepository) UpdateWalletBalance(walletId uint, newBalance float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWalletBalance", walletId, newBalance)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWalletBalance indicates an expected call of UpdateWalletBalance.
func (mr *MockIWalletRepositoryMockRecorder) UpdateWalletBalance(walletId, newBalance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWalletBalance", reflect.TypeOf((*MockIWalletRepository)(nil).UpdateWalletBalance), walletId, newBalance)
}

package business

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/util/encrypt"
	"github.com/TulioGuaraldoB/q2-payer-challenge/util/formatter"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository"
)

type IUserBusiness interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(userId uint) (*model.User, error)
	GetUserByEmail(userEmail string) (*model.User, error)
	GetUserByCpf(userCpf string) (*model.User, error)
	GetUserByCredentials(userCredentials *dto.UserCredentials) (*model.User, error)
	CreateUser(user *model.User) error
	DeleteUser(userId uint) error
}

type userBusiness struct {
	userRepository repository.IUserRepository
}

func NewUserBusiness(userRepository repository.IUserRepository) IUserBusiness {
	return &userBusiness{
		userRepository: userRepository,
	}
}

func (b *userBusiness) GetAllUsers() ([]model.User, error) {
	return b.userRepository.GetAllUsers()
}

func (b *userBusiness) GetUserById(userId uint) (*model.User, error) {
	return b.userRepository.GetUserById(userId)
}

func (b *userBusiness) GetUserByEmail(userEmail string) (*model.User, error) {
	return b.userRepository.GetUserByEmail(userEmail)
}

func (b *userBusiness) GetUserByCpf(userCpf string) (*model.User, error) {
	return b.userRepository.GetUserByCpf(userCpf)
}

func (b *userBusiness) GetUserByCredentials(userCredentials *dto.UserCredentials) (*model.User, error) {
	hashedPassword := encrypt.HashToSHA256(userCredentials.Password)
	userCredentials.Password = hashedPassword

	return b.userRepository.GetUserByCredentials(userCredentials.Email, userCredentials.Password)
}

func (b *userBusiness) CreateUser(user *model.User) error {
	unmaskedCPF := formatter.RemoveMaskFromDocument(user.CPF)
	unmaskedCNPJ := formatter.RemoveMaskFromDocument(user.CNPJ)
	user.CPF = unmaskedCPF
	user.CNPJ = unmaskedCNPJ

	hashedPassword := encrypt.HashToSHA256(user.Password)
	user.Password = hashedPassword

	return b.userRepository.CreateUser(user)
}

func (b *userBusiness) DeleteUser(userId uint) error {
	return b.userRepository.DeleteUser(userId)
}

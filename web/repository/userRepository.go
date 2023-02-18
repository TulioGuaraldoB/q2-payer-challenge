package repository

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(userId uint) (*model.User, error)
	GetUserByEmail(userEmail string) (*model.User, error)
	GetUserByCpf(userCpf string) (*model.User, error)
	GetUserByCredentials(email, password string) (*model.User, error)
	CreateUser(user *model.User) error
	DeleteUser(userId uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUserById(userId uint) (*model.User, error) {
	user := new(model.User)
	if err := r.db.First(&user, &userId).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(userEmail string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Where("email = ?", userEmail).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByCpf(userCpf string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Where("cpf = ?", userCpf).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByCredentials(email, password string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) DeleteUser(userId uint) error {
	user := new(model.User)
	return r.db.Where("id = ?", userId).Delete(&user).Error
}

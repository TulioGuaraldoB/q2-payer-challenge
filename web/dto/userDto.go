package dto

import "github.com/TulioGuaraldoB/q2-payer-challenge/web/model"

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	CPF      string `json:"cpf"`
	CNPJ     string `json:"cnpj"`
	Email    string `json:"email"`
	UserType int64  `json:"user_type"`
}

type UserRequest struct {
	FullName string `json:"full_name"`
	CPF      string `json:"cpf"`
	CNPJ     string `json:"cnpj"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType int64  `json:"user_type"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ParseUserToResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		CPF:      user.CPF,
		CNPJ:     user.CNPJ,
		Email:    user.Email,
		UserType: user.UserType,
	}
}

func ParseRequestToUser(userRequest *UserRequest) *model.User {
	return &model.User{
		FullName: userRequest.FullName,
		CPF:      userRequest.CPF,
		CNPJ:     userRequest.CNPJ,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		UserType: userRequest.UserType,
	}
}

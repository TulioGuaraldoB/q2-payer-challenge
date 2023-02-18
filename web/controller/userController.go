package controller

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/q2-payer-challenge/web/business"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IUserController interface {
	GetAllUsers(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	GetUserByEmail(ctx *fiber.Ctx) error
	GetUserByCpf(ctx *fiber.Ctx) error
	GetUserByCredentials(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}

type userController struct {
	userBusiness business.IUserBusiness
}

func NewUserController(userBusiness business.IUserBusiness) IUserController {
	return &userController{
		userBusiness: userBusiness,
	}
}

func (c *userController) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := c.userBusiness.GetAllUsers()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	if len(users) <= 0 {
		return ctx.Status(http.StatusNotFound).
			JSON(fiber.Map{"error": "data not found"})
	}

	usersResponse := make([]dto.UserResponse, 0)
	for _, user := range users {
		userResponse := dto.ParseUserToResponse(&user)
		usersResponse = append(usersResponse, *userResponse)
	}

	return ctx.Status(http.StatusOK).JSON(usersResponse)
}

func (c *userController) GetUserById(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := c.userBusiness.GetUserById(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := dto.ParseUserToResponse(user)
	return ctx.Status(http.StatusOK).JSON(userResponse)
}

func (c *userController) GetUserByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")

	user, err := c.userBusiness.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := dto.ParseUserToResponse(user)
	return ctx.Status(http.StatusOK).JSON(userResponse)
}

func (c *userController) GetUserByCpf(ctx *fiber.Ctx) error {
	cpf := ctx.Params("cpf")

	user, err := c.userBusiness.GetUserByCpf(cpf)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := dto.ParseUserToResponse(user)
	return ctx.Status(http.StatusOK).JSON(userResponse)
}

func (c *userController) GetUserByCredentials(ctx *fiber.Ctx) error {
	credentialsRequest := new(dto.UserCredentials)
	if err := ctx.BodyParser(credentialsRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := c.userBusiness.GetUserByCredentials(credentialsRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := dto.ParseUserToResponse(user)
	return ctx.Status(http.StatusOK).JSON(userResponse)
}

func (c *userController) CreateUser(ctx *fiber.Ctx) error {
	userRequest := new(dto.UserRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	user := dto.ParseRequestToUser(userRequest)
	if err := c.userBusiness.CreateUser(user); err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user created successfully!",
		"user":    userRequest,
	})
}

func (c *userController) DeleteUser(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.userBusiness.DeleteUser(uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user deleted successfully!",
	})
}

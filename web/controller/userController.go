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

// GetAllUsers godoc
// @Summary      Get All Users
// @Description  Get All Users
// @Tags         Get All Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  []dto.UserResponse
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /api/v1/user [get]
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

// GetUserById godoc
// @Summary      Get User By Id
// @Description  Get User By Id
// @Tags         Get User By Id
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.UserResponse
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /api/v1/user/:id [get]
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

// GetUserByEmail godoc
// @Summary      Get User By Email
// @Description  Get User By Email
// @Tags         Get User By Email
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.UserResponse
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /api/v1/user/email/:email [get]
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

// GetUserByCpf godoc
// @Summary      Get User By Cpf
// @Description  Get User By Cpf
// @Tags         Get User By Cpf
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.UserResponse
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /api/v1/user/cpf/:cpf [get]
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

// GetUserByCredentials godoc
// @Summary      Get User By Credentials
// @Description  Get User By Credentials
// @Tags         Get User By Credentials
// @Param        Credentials	body	dto.UserCredentials  true  "User Credentials"
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.UserResponse
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /api/v1/user/login [post]
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

// CreateUser godoc
// @Summary      Create User
// @Description  Create User
// @Tags         Create User
// @Param        User	body	dto.UserRequest  true  "User Request"
// @Accept       json
// @Produce      json
// @Success      201  {object}  dto.UserResponse
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /api/v1/user [post]
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

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully!",
		"user":    userRequest,
	})
}

// DeleteUser godoc
// @Summary      Delete User
// @Description  Delete User
// @Tags         Delete User
// @Accept       json
// @Produce      json
// @Success      200  {string}  "user deleted successfully!"
// @Failure      404  {object}  nil
// @Failure      500  {object}  nil
// @Router       /api/v1/user/:id [delete]
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

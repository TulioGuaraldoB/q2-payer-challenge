package controller

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/q2-payer-challenge/web/business"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ITransactionController interface {
	GetTransactionById(ctx *fiber.Ctx) error
}

type transactionController struct {
	transactionBusiness business.ITransactionBusiness
}

func NewTransactionController(transactionBusiness business.ITransactionBusiness) ITransactionController {
	return &transactionController{
		transactionBusiness: transactionBusiness,
	}
}

func (c *transactionController) GetTransactionById(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	transaction, err := c.transactionBusiness.GetTransactionById(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	transactionResponse := dto.ParseTransactionToResponse(transaction)
	return ctx.Status(http.StatusOK).JSON(transactionResponse)
}

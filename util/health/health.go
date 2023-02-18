package health

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).
		JSON("I heard that Tulio passed in the Mid-Level Back-End challenge! ( ͡¬ ͜ʖ ͡¬)")
}

package exception

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		var ve *ValidationError
		if errors.As(err, &ve) {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"success": false,
				"message": ve.Errs,
			})
		}

		return ctx.Status(code).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
}

package exception

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// NewErrorHandler returns a custom error handler for a Fiber application.
// This handler processes different types of errors and responds with appropriate
// HTTP status codes and JSON messages.
//
// The function returned handles errors as follows:
//  1. If the error is of type *fiber.Error, it sets the response status code
//     to the code specified in the fiber.Error.
//  2. If the error is of type *ValidationError, it responds with status code
//     422 (Unprocessable Entity) and a JSON object containing the validation errors.
//  3. For all other errors, it responds with status code 500 (Internal Server Error)
//     and a JSON object containing the error message.
//
// The JSON response structure for handled errors is:
//
//	{
//	    "success": false,
//	    "message": "<error message or validation errors>"
//	}
//
// Example:
//
//	app := fiber.New(fiber.Config{
//	  ErrorHandler: exception.NewErrorHandler(),
//	})
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

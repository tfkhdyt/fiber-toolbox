package exception

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// NewInternalServerError creates a new fiber.Error with a status code of 500 (Internal Server Error).
// If an error is provided in the errs variadic parameter, it logs the error and includes its message as the cause
// in the returned error message.
//
// Parameters:
//   - message: A string representing the error message.
//   - errs: Optional variadic parameter for including a cause of the error.
//
// Returns:
//   - error: A new fiber.Error with status code 500 and the provided message (including the cause if errs is provided).
func NewInternalServerError(message string, errs ...error) error {
	if len(errs) > 0 {
		log.Errorw(message, "cause", errs[0])
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("%v. Cause: %v", message, errs[0].Error()))
	}
	return fiber.NewError(fiber.StatusInternalServerError, message)
}

// NewNotFoundError creates a new fiber.Error with a status code of 404 (Not Found).
// If an error is provided in the errs variadic parameter, it logs the warning with the error.
//
// Parameters:
//   - message: A string representing the error message.
//   - errs: Optional variadic parameter for including a cause of the error.
//
// Returns:
//   - error: A new fiber.Error with status code 404 and the provided message.
func NewNotFoundError(message string, errs ...error) error {
	if len(errs) > 0 {
		log.Warnw(message, "cause", errs[0])
	}
	return fiber.NewError(fiber.StatusNotFound, message)
}

// NewBadRequestError creates a new fiber.Error with a status code of 400 (Bad Request).
// If an error is provided in the errs variadic parameter, it logs the warning and includes its message as the cause
// in the returned error message.
//
// Parameters:
//   - message: A string representing the error message.
//   - errs: Optional variadic parameter for including a cause of the error.
//
// Returns:
//   - error: A new fiber.Error with status code 400 and the provided message (including the cause if errs is provided).
func NewBadRequestError(message string, errs ...error) error {
	if len(errs) > 0 {
		log.Warnw(message, "cause", errs[0])
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("%v. Cause: %v", message, errs[0].Error()))
	}
	return fiber.NewError(fiber.StatusBadRequest, message)
}

// NewUnauthenticatedError creates a new fiber.Error with a status code of 401 (Unauthorized).
// If an error is provided in the errs variadic parameter, it logs the warning with the error.
//
// Parameters:
//   - message: A string representing the error message.
//   - errs: Optional variadic parameter for including a cause of the error.
//
// Returns:
//   - error: A new fiber.Error with status code 401 and the provided message.
func NewUnauthenticatedError(message string, errs ...error) error {
	if len(errs) > 0 {
		log.Warnw(message, "cause", errs[0])
	}
	return fiber.NewError(fiber.StatusUnauthorized, message)
}

// NewUnauthorizedError creates a new fiber.Error with a status code of 403 (Forbidden).
// If an error is provided in the errs variadic parameter, it logs the warning with the error.
//
// Parameters:
//   - message: A string representing the error message.
//   - errs: Optional variadic parameter for including a cause of the error.
//
// Returns:
//   - error: A new fiber.Error with status code 403 and the provided message.
func NewUnauthorizedError(message string, errs ...error) error {
	if len(errs) > 0 {
		log.Warnw(message, "cause", errs[0])
	}
	return fiber.NewError(fiber.StatusForbidden, message)
}

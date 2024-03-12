package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func NewInternalServerError(message string, errs ...error) error {
	log.Errorw(message, "err", errs[0])
	return fiber.NewError(fiber.StatusInternalServerError, message)
}

func NewNotFoundError(message string, errs ...error) error {
	log.Warnw(message, "warn", errs[0])
	return fiber.NewError(fiber.StatusNotFound, message)
}

func NewBadRequestError(message string, errs ...error) error {
	log.Warnw(message, "warn", errs[0])
	return fiber.NewError(fiber.StatusBadRequest, message)
}

func NewUnauthenticatedError(message string, errs ...error) error {
	log.Warnw(message, "warn", errs[0])
	return fiber.NewError(fiber.StatusUnauthorized, message)
}

func NewUnauthorizedError(message string, errs ...error) error {
	return fiber.NewError(fiber.StatusForbidden, message)
}

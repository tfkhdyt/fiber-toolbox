package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

func ParsePayloadFromHeaders[T jwt.Claims](c *fiber.Ctx) (*T, error) {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil, exception.NewBadRequestError("failed to validate token")
	}

	claims, ok := token.Claims.(T)
	if !ok {
		return nil, exception.NewBadRequestError("failed to validate claims")
	}

	return &claims, nil
}

func ParsePayload[T jwt.Claims](tokenString string, jwtKey string) (*T, error) {
	clm := new(T)

	token, err := jwt.ParseWithClaims(
		tokenString,
		*clm,
		func(token *jwt.Token) (any, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, exception.NewBadRequestError(
			"failed to parse jwt payload",
			err,
		)
	}

	claims, ok := token.Claims.(T)
	if !ok {
		return nil, exception.NewBadRequestError("failed to validate claims")
	}

	return &claims, nil
}

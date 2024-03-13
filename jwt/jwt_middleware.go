package jwt

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

func JwtMiddleware[T jwt.Claims](jwtKey string) func(c *fiber.Ctx) error {
	clm := new(T)

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtKey)},
		Claims:     *clm,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return exception.NewUnauthenticatedError(err.Error())
		},
	})
}

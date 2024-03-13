package jwt

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

func JwtMiddleware[T StructClaims]() func(c *fiber.Ctx) error {
	clm := new(T)

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_ACCESS_KEY"))},
		Claims:     *clm,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return exception.NewUnauthenticatedError(err.Error())
		},
	})
}

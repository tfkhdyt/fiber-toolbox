package jwt

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

// JwtMiddleware returns a JWT middleware function for Fiber that validates incoming requests' JWT tokens.
// It uses the provided StructClaims type for the token's claims.
//
// Required Env:
//   - JWT_ACCESS_KEY = Secret key for signing jwt access token
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

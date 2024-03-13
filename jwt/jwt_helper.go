package jwt

import (
	"os"
	"time"

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

type JwtType uint

const (
	Access JwtType = iota
	Refresh
)

func GenerateJWT(claims *jwt.RegisteredClaims, jwtType JwtType) (string, error) {
	var exp time.Time
	var jwtKey string

	switch jwtType {
	case Access:
		exp = time.Now().Add(24 * time.Hour * 7)
		jwtKey = os.Getenv("JWT_ACCESS_KEY")
	case Refresh:
		exp = time.Now().Add(24 * time.Hour * 30)
		jwtKey = os.Getenv("JWT_REFRESH_KEY")
	default:
		return "", exception.NewBadRequestError("invalid JWT type")
	}

	claims.ExpiresAt = jwt.NewNumericDate(exp)
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.NotBefore = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", exception.NewInternalServerError(
			"failed to sign jwt token",
			err,
		)
	}

	return t, nil
}
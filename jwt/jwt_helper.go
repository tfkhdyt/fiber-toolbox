package jwt

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

// ParsePayloadFromHeaders extracts JWT claims from the request context and validates them.
// It expects the JWT token to be stored in the request's context under the "user" key.
func ParsePayloadFromHeaders[T StructClaims](c *fiber.Ctx, claims T) (T, error) {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return claims, exception.NewBadRequestError("failed to validate token")
	}

	clm, ok := token.Claims.(T)
	if !ok {
		return claims, exception.NewBadRequestError("failed to validate claims")
	}

	return clm, nil
}

// ParsePayload parses a JWT token string and validates its claims using the provided JWT key.
// It returns the claims if the token is valid, or an error if not.
func ParsePayload[T StructClaims](tokenString string, jwtKey string, claims T) (T, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return claims, exception.NewBadRequestError(
			"failed to parse jwt payload",
			err,
		)
	}

	claims, ok := token.Claims.(T)
	if !ok {
		return claims, exception.NewBadRequestError("failed to validate claims")
	}

	return claims, nil
}

// JwtType defines the type of JWT, either Access or Refresh.
type JwtType uint

const (
	Access  JwtType = iota // Access token type
	Refresh                // Refresh token type
)

// StructClaims is an interface that extends jwt.Claims with a method to set the expiration time.
type StructClaims interface {
	jwt.Claims
	SetExp(time.Time)
}

// GenerateJWT generates a JWT token with specified claims and type (Access or Refresh).
// The expiration time and signing key are determined based on the token type.
//
// Required Env:
//   - JWT_ACCESS_KEY = Secret key for signing jwt access token
//   - JWT_REFRESH_KEY = Secret key for signing jwt refresh token
func GenerateJWT[T StructClaims](claims T, jwtType JwtType) (string, error) {
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

	claims.SetExp(exp)

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

package hash

import (
	"github.com/matthewhartstonge/argon2"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

// Global variable for Argon2 configuration
var argon argon2.Config

// init function initializes the Argon2 configuration with default settings
func init() {
	argon = argon2.DefaultConfig()
}

// HashPassword hashes a given password using Argon2 and returns the hashed password
// If hashing fails, it returns an internal server error
func HashPassword(password string) (string, error) {
	hashedPassword, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", exception.NewInternalServerError("failed to hash password")
	}

	return string(hashedPassword), nil
}

// VerifyPassword verifies a given password against a hashed password using Argon2
// If verification fails, it returns an appropriate error (internal server error or bad request error)
func VerifyPassword(password string, hashedPwd string) error {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hashedPwd))
	if err != nil {
		return exception.NewInternalServerError("failed to verify password")
	}

	if !ok {
		return exception.NewBadRequestError("invalid password")
	}

	return nil
}

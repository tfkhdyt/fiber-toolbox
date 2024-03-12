package service

import (
	"github.com/matthewhartstonge/argon2"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

var argon argon2.Config

func init() {
	argon = argon2.DefaultConfig()
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", exception.NewInternalServerError("failed to hash password")
	}

	return string(hashedPassword), nil
}

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

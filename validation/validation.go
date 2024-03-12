package helper

import (
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New(validator.WithRequiredStructEnabled())
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalf("error(validator): %v", err)
	}
}

func validateStruct(payload any) []string {
	errs := []string{}

	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, e := range validationErrors {
			errs = append(errs, e.Translate(trans))
		}

		return errs
	}

	return nil
}

func ValidateBody(c *fiber.Ctx, payload any) error {
	if err := c.BodyParser(payload); err != nil {
		log.Println("error(body-parser):", err)
		return fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"invalid request body",
		)
	}

	if err := validateStruct(payload); err != nil {
		return exception.NewValidationError(err)
	}

	return nil
}

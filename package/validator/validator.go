package validator

import (
	"errors"
	"typer/package/exceptions"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validatable interface {
	ErrorMessages() map[string]string
}

func Validate(ctx *fiber.Ctx, dto interface{}) error{
	if err := ctx.BodyParser(dto); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return err
	}

	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		errorBag := make(map[string]string)

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			customMessages := make(map[string]string)

			if v, ok := dto.(Validatable); ok {
				customMessages = v.ErrorMessages()
			}

			for _, fieldErr := range validationErrors {
				field := fieldErr.Field()
				tag := fieldErr.Tag()
				key := field + "." + tag

				if msg, exists := customMessages[key]; exists {
					errorBag[field] = msg
				} else {
					errorBag[field] = field + " is invalid"
				}
			}
		}

		return exceptions.ValidationError{
			Message: "Validation failed",
			Errors:  errorBag,
		}
	}

	return nil
}
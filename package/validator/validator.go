package validator

import (
	"errors"
	"strings"
	"typer/package/exceptions"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validatable interface {
	ErrorMessages() map[string]string
}

func Validate(ctx *fiber.Ctx, dto any) error {
	if err := ctx.BodyParser(dto); err != nil {
		errorMsg := err.Error()
		if strings.Contains(errorMsg, "invalid character") || strings.Contains(errorMsg, "unexpected end") {
			return &exceptions.ValidationError{
				Message: "Invalid JSON format",
				Errors: map[string]string{
					"body": "JSON syntax error: " + errorMsg + ". Make sure all strings are quoted with double quotes.",
				},
			}
		}
		return &exceptions.ValidationError{
			Message: "Invalid request body",
			Errors: map[string]string{
				"body": "Failed to parse request body: " + errorMsg,
			},
		}
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

				var key string
				if fieldErr.Namespace() != fieldErr.Field() {
					key = field + "." + tag
				} else {
					key = field + "." + tag
				}

				if msg, exists := customMessages[key]; exists {
					errorBag[field] = msg
				} else {
					errorBag[field] = generateDefaultMessage(field, tag, fieldErr)
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

func generateDefaultMessage(field, tag string, fieldErr validator.FieldError) string {
	switch tag {
	case "required":
		return field + " is required"
	case "min":
		return field + " must have at least " + fieldErr.Param() + " items"
	case "max":
		return field + " cannot have more than " + fieldErr.Param() + " items"
	case "dive":
		return "Invalid items in " + field
	default:
		return field + " is invalid"
	}
}

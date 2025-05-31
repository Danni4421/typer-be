package controllers

import (
	"reflect"
	"typer/app/dto"
	"typer/app/services"
	"typer/package/exceptions"
	"typer/package/utils"
	"typer/package/validator"

	"github.com/gofiber/fiber/v2"
)

type TypingController struct {
	LanguageService *services.LanguageService
	TypeService     *services.TypingService
}

func (t *TypingController) StoreTypingTestLog(ctx *fiber.Ctx) error {
	userID, ok := utils.ParseUserIDFromLocals(ctx)

	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "You are not authorized, please authenticate first",
		})
	}

	langCode := ctx.Params("code")

	if langCode == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Language code is required",
		})
	}

	typingLogDto := new(dto.StoreUserLogDto)

	if err := validator.Validate(ctx, typingLogDto); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ValidationError{}) {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.(exceptions.ValidationError).ToMap())
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	lang, err := t.LanguageService.GetLanguageByCode(langCode)

	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(&exceptions.ClientError{}) {
			return ctx.Status(err.(*exceptions.ClientError).Code).JSON(fiber.Map{
				"status":  "error",
				"message": err.(*exceptions.ClientError).Message,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	if err := t.TypeService.StoreTypingTestLog(uint(userID), lang.ID, &typingLogDto.Calculation); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(&exceptions.ClientError{}) {
			return ctx.Status(err.(*exceptions.ClientError).Code).JSON(fiber.Map{
				"status":  "error",
				"message": err.(*exceptions.ClientError).Message,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Typing test log stored successfully",
	})
}

func (t *TypingController) CalculateTypingTestScore(ctx *fiber.Ctx) error {
	typingCalculationDto := new(dto.TypingCalculationDto)

	if err := validator.Validate(ctx, typingCalculationDto); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ValidationError{}) {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.(exceptions.ValidationError).ToMap())
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	calculationResult, err := t.TypeService.CalculateWPM(typingCalculationDto.Text, typingCalculationDto.FailedText)

	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(&exceptions.ClientError{}) {
			return ctx.Status(err.(*exceptions.ClientError).Code).JSON(fiber.Map{
				"status":  "error",
				"message": err.(*exceptions.ClientError).Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Typing test score calculated successfully",
		"data":    calculationResult,
	})
}

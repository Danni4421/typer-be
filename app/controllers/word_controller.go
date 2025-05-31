package controllers

import (
	"fmt"
	"strconv"
	"typer/app/dto"
	"typer/app/services"
	"typer/package/exceptions"
	"typer/package/validator"

	"github.com/gofiber/fiber/v2"
)

type WordController struct {
	LanguageService *services.LanguageService
	WordService     *services.WordService
}

func (ctrl *WordController) StoreWords(ctx *fiber.Ctx) error {
	languageCode := ctx.Params("code")

	if languageCode == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid code parameter",
		})
	}

	storeWordsDto := new(dto.StoreWords)
	if err := validator.Validate(ctx, storeWordsDto); err != nil {
		fmt.Println("Validation error:", err)
		if validationErr, ok := err.(exceptions.ValidationError); ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(validationErr.ToMap())
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	language, err := ctrl.LanguageService.GetLanguageByCode(string(languageCode))
	if err != nil {
		if clientErr, ok := err.(*exceptions.ClientError); ok {
			return ctx.Status(clientErr.Code).JSON(fiber.Map{
				"status":  "error",
				"message": clientErr.Message,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}
	if err := ctrl.WordService.StoreWords(storeWordsDto.Words, language.ID); err != nil {
		if clientErr, ok := err.(*exceptions.ClientError); ok {
			return ctx.Status(clientErr.Code).JSON(fiber.Map{
				"status":  "error",
				"message": clientErr.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Words stored successfully",
		"data":    storeWordsDto.Words,
	})
}

func (ctrl *WordController) GetWordsByLanguage(ctx *fiber.Ctx) error {
	languageCode := ctx.Params("code")

	if languageCode == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid language code",
		})
	}

	language, err := ctrl.LanguageService.GetLanguageByCode(languageCode)
	if err != nil {
		if clientErr, ok := err.(*exceptions.ClientError); ok {
			return ctx.Status(clientErr.Code).JSON(fiber.Map{
				"status":  "error",
				"message": clientErr.Message,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	words, err := ctrl.WordService.GetWordsByLanguage(language.ID)

	if err != nil {
		if clientErr, ok := err.(*exceptions.ClientError); ok {
			return ctx.Status(clientErr.Code).JSON(fiber.Map{
				"status":  "error",
				"message": clientErr.Message,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Words retrieved successfully",
		"data":    words,
	})
}

func (ctrl *WordController) GetRandomWords(ctx *fiber.Ctx) error {
	languageCode := ctx.Params("code")
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))

	if err != nil || limit <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid limit parameter",
			"errors":  err.Error(),
		})
	}

	language, err := ctrl.LanguageService.GetLanguageByCode(languageCode)
	if err != nil {
		if clientErr, ok := err.(*exceptions.ClientError); ok {
			return ctx.Status(clientErr.Code).JSON(fiber.Map{
				"status":  "error",
				"message": clientErr.Message,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	randomWords, err := ctrl.WordService.GetRandomWords(language.ID, limit)

	if err != nil {
		if clientErr, ok := err.(*exceptions.ClientError); ok {
			return ctx.Status(clientErr.Code).JSON(fiber.Map{
				"status":  "error",
				"message": clientErr.Message,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Random words retrieved successfully",
		"data":    randomWords,
	})
}

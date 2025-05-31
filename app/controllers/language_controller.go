package controllers

import (
	"reflect"
	"typer/app/dto"
	"typer/app/services"
	"typer/package/exceptions"
	"typer/package/validator"

	"github.com/gofiber/fiber/v2"
)

type LanguageController struct {
	LanguageService *services.LanguageService
}

func (ctrl *LanguageController) CreateNewLanguage(ctx *fiber.Ctx) error {
	createLanguageDto := new(dto.CreateLanguage)

	if err := validator.Validate(ctx, createLanguageDto); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ValidationError{}) {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.(exceptions.ValidationError).ToMap())
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	if err := ctrl.LanguageService.CreateLanguage(createLanguageDto.Name, createLanguageDto.Code); err != nil {
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

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Language created successfully",
		"data": fiber.Map{
			"code": createLanguageDto.Code,
			"name": createLanguageDto.Name,
		},
	})
}

func (ctrl *LanguageController) GetAllLanguages(ctx *fiber.Ctx) error {
	languages, err := ctrl.LanguageService.GetAllLanguages()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve languages",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Languages retrieved successfully",
		"data":    languages,
	})
}

func (ctrl *LanguageController) GetLanguageByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid language name",
		})
	}

	language, err := ctrl.LanguageService.GetLanguageByName(name)
	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(&exceptions.ClientError{}) {
			return ctx.Status(err.(*exceptions.ClientError).Code).JSON(err.(*exceptions.ClientError).ToMap())
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Language retrieved successfully",
		"data":    language,
	})
}

func (ctrl *LanguageController) DeleteLanguageByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	if code == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid language code",
		})
	}

	if err := ctrl.LanguageService.DeleteLanguageByCode(code); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(&exceptions.ClientError{}) {
			return ctx.Status(err.(*exceptions.ClientError).Code).JSON(err.(*exceptions.ClientError).ToMap())
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Language deleted successfully",
	})
}

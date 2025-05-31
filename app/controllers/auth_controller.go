package controllers

import (
	"fmt"
	"reflect"
	"typer/app/dto"
	"typer/app/services"
	"typer/package/exceptions"
	"typer/package/validator"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService *services.AuthService
	JWTService  *services.JWTService
}

func (ctrl *AuthController) Login(ctx *fiber.Ctx) error {
	loginDto := new(dto.LoginUser)

	if err := validator.Validate(ctx, loginDto); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ValidationError{}) {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.(exceptions.ValidationError).ToMap())
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	userID, err := ctrl.AuthService.ValidateCredentials(loginDto.Email, loginDto.Password)

	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ClientError{}) {
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
	tokens, err := ctrl.JWTService.GenerateTokens(fmt.Sprintf("%d", userID))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to generate tokens",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login successful",
		"data": fiber.Map{
			"token":         tokens.Token,
			"refresh_token": tokens.RefreshToken,
		},
	})
}

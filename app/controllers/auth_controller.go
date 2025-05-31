package controllers

import (
	"fmt"
	"reflect"
	"typer/app/dto"
	"typer/app/services"
	"typer/package/exceptions"
	"typer/package/utils"
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

func (ctrl *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	refreshTokenDto := new(dto.RefreshTokenRequest)

	if err := validator.Validate(ctx, refreshTokenDto); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ValidationError{}) {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.(exceptions.ValidationError).ToMap())
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	userID, err := ctrl.JWTService.ValidateRefreshToken(refreshTokenDto.RefreshToken)
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

	tokens, err := ctrl.JWTService.RenewTokens(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to generate tokens",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Token refreshed successfully",
		"data": fiber.Map{
			"token":         tokens.Token,
			"refresh_token": tokens.RefreshToken,
		},
	})
}

func (ctrl *AuthController) Logout(ctx *fiber.Ctx) error {
	userID, ok := utils.ParseUserIDFromLocals(ctx)

	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized access",
		})
	}

	err := ctrl.AuthService.Logout(userID)
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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Logout successful",
	})
}

package controllers

import (
	"reflect"
	"typer/app/dto"
	"typer/app/models"
	"typer/app/services"
	"typer/package/exceptions"
	"typer/package/validator"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func (u *UserController) RegisterNewUser(c *fiber.Ctx) error {
	registerUserDto := new(dto.RegisterUser)

	if err := validator.Validate(c, registerUserDto); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ValidationError{}) {
			return c.Status(fiber.StatusBadRequest).JSON(err.(exceptions.ValidationError).ToMap())
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	createdUser, err := u.UserService.CreateUser(&models.User{
		Username: registerUserDto.Username,
		Name:     registerUserDto.Name,
		Email:    registerUserDto.Email,
		Password: registerUserDto.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully",
		"data":    createdUser,
	})
}

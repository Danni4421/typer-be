package controllers

import (
	"reflect"
	"typer/app/dto"
	"typer/app/models"
	"typer/app/services"
	"typer/package/exceptions"
	"typer/package/utils"
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

func (u *UserController) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user username",
		})
	}

	user, err := u.UserService.GetUserByUsername(username)

	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ClientError{}) {
			return c.Status(err.(*exceptions.ClientError).Code).JSON(err.(*exceptions.ClientError).ToMap())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (u *UserController) GetCurrentUser(c *fiber.Ctx) error {
	userID, ok := utils.ParseUserIDFromLocals(c)

	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user ID",
		})
	}

	user, err := u.UserService.GetUserByID(userID)

	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(exceptions.ClientError{}) {
			return c.Status(err.(*exceptions.ClientError).Code).JSON(err.(*exceptions.ClientError).ToMap())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal server error",
			"errors":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}

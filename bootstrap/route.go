package bootstrap

import (
	"typer/app/controllers"
	"typer/app/services"
	"typer/platform/database"

	"github.com/gofiber/fiber/v2"
)

var userController *controllers.UserController

func SetupControllers() {
	// Services
	userService := services.UserService{
		DB: database.DB,
	}

	// Controllers
	userController = &controllers.UserController{
		UserService: &userService,
	}
}

func bindRoutes(app *fiber.App) {
	app.Group("/api")

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "pong",
		})
	})

	app.Post("/users", userController.RegisterNewUser)
}

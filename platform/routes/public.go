package routes

import (
	typer "typer/app"

	"github.com/gofiber/fiber/v2"
)

func BindPublicRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "pong"})
	})
	api.Post("/auth/login", typer.AuthController.Login)
	api.Post("/auth/register", typer.UserController.RegisterNewUser)
	api.Get("/users/:username", typer.UserController.GetUserByUsername)
}

package routes

import (
	typer "typer/app"

	"github.com/gofiber/fiber/v2"
)

func BindPublicRoutes(app *fiber.App) {
	api := app.Group("/pub/api")

	api.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "pong"})
	})

	api.Get("/token/csrf", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"csrf_token": ctx.Locals("token"),
		})
	})

	// Auth
	api.Post("/auth/login", typer.AuthController.Login)
	api.Put("/auth/refresh", typer.AuthController.RefreshToken)
	api.Post("/auth/register", typer.UserController.RegisterNewUser)

	// Users
	api.Get("/users/:username", typer.UserController.GetUserByUsername)

	// Words
	api.Get("/languages", typer.LanguageController.GetAllLanguages)
	api.Get("/languages/:name", typer.LanguageController.GetLanguageByName)
	api.Get("/languages/:code/words/random", typer.WordController.GetRandomWords)

	// Typing
	api.Post("/typing/calculate", typer.TypingController.CalculateTypingTestScore)
}

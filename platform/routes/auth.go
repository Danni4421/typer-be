package routes

import (
	typer "typer/app"
	"typer/platform/middleware"

	"github.com/gofiber/fiber/v2"
)

func BindAuthenticatedRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.JWTMiddleware())

	api.Get("/users/me", typer.UserController.GetCurrentUser)

	// Languages
	api.Post("/languages", typer.LanguageController.CreateNewLanguage)
	api.Delete("/languages/:code", typer.LanguageController.DeleteLanguageByCode)

	// Words
	api.Post("/languages/:code/words", typer.WordController.StoreWords)
	api.Get("/languages/:code/words", typer.WordController.GetWordsByLanguage)
}

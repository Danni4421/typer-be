package routes

import (
	typer "typer/app"
	"typer/platform/middleware"

	"github.com/gofiber/fiber/v2"
)

func BindAuthenticatedRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.JWTMiddleware())

	api.Get("/users/me", typer.UserController.GetCurrentUser)
}

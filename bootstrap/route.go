package bootstrap

import "github.com/gofiber/fiber/v2"

func bindRoutes(app *fiber.App) {
	app.Group("/api")

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "pong",
		})
	})
}

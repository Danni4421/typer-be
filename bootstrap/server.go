package bootstrap

import (
	"fmt"
	"typer/platform/database"

	"github.com/gofiber/fiber/v2"
)

func listenOnPanic() {
	if res := recover(); res != nil {
		fmt.Println("Recovered from panic:", res)
	}
}

func App() {
	defer listenOnPanic()

	app := fiber.New()

	database.ConnectPostgres()
	SetupControllers()

	bindRoutes(app)

	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}

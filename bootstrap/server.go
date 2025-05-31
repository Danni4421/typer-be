package bootstrap

import (
	"fmt"

	typer "typer/app"
	"typer/package/configs"
	"typer/platform/database"
	"typer/platform/middleware"
	"typer/platform/routes"

	"github.com/gofiber/fiber/v2"
)

func listenOnPanic() {
	if res := recover(); res != nil {
		fmt.Println("Recovered from panic:", res)
	}
}

func App() {
	defer listenOnPanic()

	// Load the configuration from the config file
	appConfig := configs.GetFiberConfig()

	//  Create a new Fiber app with custom configuration
	app := fiber.New(appConfig)

	// Setup the database connection and controllers
	database.ConnectPostgres()
	typer.SetupControllers()

	// Register global middleware
	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.CSRFMiddleware())
	app.Use(middleware.LoggerMiddleware())

	// Bind routes to the app
	routes.BindAuthenticatedRoutes(app)
	routes.BindPublicRoutes(app)

	// Start the server and listen on port 3000 by default
	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}

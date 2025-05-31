package app

import (
	"typer/app/controllers"
	"typer/app/services"
	"typer/platform/database"
)

var (
	UserController     *controllers.UserController
	AuthController     *controllers.AuthController
	LanguageController *controllers.LanguageController
	WordController     *controllers.WordController
	TypingController   *controllers.TypingController
)

func SetupControllers() {
	// Define services
	// The services are initialized with the database connection
	// They are then used by the controllers to perform business logic
	userService := services.UserService{
		DB: database.DB,
	}

	authService := services.AuthService{
		DB: database.DB,
	}

	jwtService := services.JWTService{
		DB: database.DB,
	}

	languageService := services.LanguageService{
		DB: database.DB,
	}

	wordService := services.WordService{
		DB: database.DB,
	}

	typingService := services.TypingService{
		DB: database.DB,
	}

	// Define controllers
	// The controller are initialized with their respective services
	// Then they are assigned to the endpoints in the bindRoutes function
	UserController = &controllers.UserController{
		UserService: &userService,
	}

	AuthController = &controllers.AuthController{
		AuthService: &authService,
		JWTService:  &jwtService,
	}

	LanguageController = &controllers.LanguageController{
		LanguageService: &languageService,
	}

	WordController = &controllers.WordController{
		LanguageService: &languageService,
		WordService:     &wordService,
	}

	TypingController = &controllers.TypingController{
		LanguageService: &languageService,
		TypeService:     &typingService,
	}
}

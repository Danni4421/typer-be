package middleware

import (
	"typer/package/exceptions"
	"typer/package/utils"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		actualToken := c.Get("Authorization")

		if actualToken == "" {
			return &exceptions.ClientError{
				Code:    401,
				Message: "You are not authorized, please authenticate first",
			}
		}

		const bearerPrefix = "Bearer "
		if len(actualToken) <= len(bearerPrefix) || actualToken[:len(bearerPrefix)] != bearerPrefix {
			return &exceptions.ClientError{
				Code:    401,
				Message: "Invalid authorization header format",
			}
		}
		authToken := actualToken[len(bearerPrefix):]

		userID, err := utils.ParseJWT(authToken)
		if err != nil {
			return &exceptions.ClientError{
				Code:    401,
				Message: "You are not authorized, please authenticate first",
			}
		}

		c.Locals("userID", userID)

		return c.Next()
	}
}

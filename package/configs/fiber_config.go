package configs

import (
	"errors"
	"log"
	"strconv"
	"time"
	"typer/package/exceptions"
	"typer/package/utils"

	"github.com/gofiber/fiber/v2"
)

func GetFiberConfig() fiber.Config {
	readTimeout, _ := strconv.Atoi(utils.GetEnv("APP_READ_TIMEOUT", "10"))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeout),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var resp any

			// Print the stack trace for debugging
			if utils.GetEnv("APP_ENV", "production") == "development" {
				log.Println(err)
			}

			var clientErr *exceptions.ClientError
			var serverErr *exceptions.ServerError

			switch {
			case errors.As(err, &clientErr):
				resp = fiber.Map{
					"status":  "error",
					"message": clientErr.Message,
				}
			case errors.As(err, &serverErr):
				resp = fiber.Map{
					"status":  "error",
					"message": serverErr.Message,
				}
			default:
				if fiberErr, ok := err.(*fiber.Error); ok {
					code = fiberErr.Code
					resp = fiber.Map{
						"status":  "error",
						"message": fiberErr.Message,
					}
				} else {
					resp = fiber.Map{
						"status":  "error",
						"message": err.Error(),
					}
				}
			}
			return c.Status(code).JSON(resp)
		},
	}
}

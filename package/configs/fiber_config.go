package configs

import (
	"log"
	"strconv"
	"time"
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

			switch e := err.(type) {
			case *fiber.Error:
				code = e.Code
				resp = fiber.Map{
					"status":  "error",
					"message": e.Message,
				}
			case interface{ ToMap() map[string]any }:
				resp = e.ToMap()
				if status, ok := resp.(map[string]any)["code"]; ok {
					if statusCode, ok := status.(int); ok {
						code = statusCode
					}
				}
				if status, ok := resp.(map[string]any)["status"]; ok {
					if statusCode, ok := status.(int); ok {
						code = statusCode
					}
				}
			default:
				resp = fiber.Map{
					"status":  "error",
					"message": err.Error(),
				}
			}
			return c.Status(code).JSON(resp)
		},
	}
}

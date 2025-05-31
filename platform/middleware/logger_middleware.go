package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerMiddleware() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} ${status} ${method} ${path} - ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "UTC",
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/health"
		},
	})
}

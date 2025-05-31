package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func CSRFMiddleware() fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:      "header:X-CSRF-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		CookieHTTPOnly: true,
		CookieSecure:   false,
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUIDv4,

		// Bypass CSRF middleware for specific routes
		// This allows public routes to be accessed without CSRF protection
		Next: func(c *fiber.Ctx) bool {
			if c.Path() == "/api/auth/login" || c.Path() == "/api/auth/register" {
				return true
			}
			if c.Method() == "GET" {
				return true
			}
			return false
		},
	})
}

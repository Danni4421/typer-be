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
		CookieSecure:   true,
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUIDv4,
	})
}

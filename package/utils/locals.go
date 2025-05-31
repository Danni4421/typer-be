package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseUserIDFromLocals(c *fiber.Ctx) (int, bool) {
	userIDValue := c.Locals("userID")

	if userIDValue == nil {
		return 0, false
	}

	switch v := userIDValue.(type) {
	case float64:
		return int(v), true
	case int:
		return v, true
	case string:
		id, err := strconv.Atoi(v)
		if err != nil {
			return 0, false
		}
		return id, true
	default:
		return 0, false
	}
}

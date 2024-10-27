package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FakeUser(c *fiber.Ctx) error {
	username := c.Cookies("test")
	id := fmt.Sprintf("%d", time.Now().Unix())

	if username == "" {
		cookie := new(fiber.Cookie)
		cookie.Name = "test"
		cookie.Value = id
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.HTTPOnly = true
		c.Cookie(cookie)
	}

	return c.Next()
}

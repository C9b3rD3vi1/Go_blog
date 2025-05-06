// middleware/auth.go
package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireAdminAuth(c *fiber.Ctx) error {
	admin := c.Locals("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}
	return c.Next()
}

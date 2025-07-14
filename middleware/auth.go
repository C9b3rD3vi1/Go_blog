// middleware/auth.go
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New()


func RequireAdminAuth(c *fiber.Ctx) error {
	admin := c.Locals("admin")
	if admin == nil {
		// If admin is not authenticated, redirect to login page
		c.Status(401)
		return c.Redirect("/admin/login")
	}

	// check if the user is admin
	if !admin.(*models.User).IsAdmin {
		c.Status(403)
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Next()

}

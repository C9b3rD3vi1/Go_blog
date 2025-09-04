// middleware/auth.go
package middleware

import (
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
	//		"github.com/gofiber/fiber/v2/middleware/session"
	//	 "github.com/gofiber/fiber/v2/utils"
)

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

// LogoutUser handles user logout
func LogoutUser(c *fiber.Ctx) error {
    sess, err := config.Store.Get(c)
    if err != nil {
        return err
    }
    if err := sess.Destroy(); err != nil {
        return err
    }
    return c.Redirect("/admin/login")
}

// middleware/auth.go
package middleware

import (
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/gofiber/fiber/v2"
	//		"github.com/gofiber/fiber/v2/middleware/session"
	//	 "github.com/gofiber/fiber/v2/utils"
)

func RequireAdminAuth(c *fiber.Ctx) error {
    // Get session
    sess, err := config.Store.Get(c)
    if err != nil {
        return c.Redirect("/admin/login")
    }

    adminID := sess.Get("admin_id")
    if adminID == nil {
        return c.Redirect("/admin/login")
    }

    // Fetch user from DB
    var admin models.User
    if err := database.DB.First(&admin, adminID).Error; err != nil {
        return c.Redirect("/admin/login")
    }

    // Check if user is actually an admin
    if !admin.IsAdmin {
        return c.SendStatus(fiber.StatusForbidden) // 403
    }

    // Store in context for handlers to use
    c.Locals("admin", &admin)

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

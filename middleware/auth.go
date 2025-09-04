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

    // ✅ match key with CreateUserSession
    userID := sess.Get("user_id")
    if userID == nil {
        return c.Redirect("/admin/login")
    }

    // Fetch user from DB
    var user models.User
    if err := database.DB.First(&user, userID).Error; err != nil {
        return c.Redirect("/admin/login")
    }

    // Check if user is actually an admin
    if !user.IsAdmin {
        return c.SendStatus(fiber.StatusForbidden) // 403
    }

    // ✅ store back in context for downstream handlers
    c.Locals("user", &user)

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

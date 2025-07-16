// middleware/auth.go
package middleware

import (
	"time"

	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New(session.Config{
	Expiration:        time.Hour * 24,
	Storage:           nil,
	KeyLookup:         "",
	CookieDomain:      "",
	CookiePath:        "",
	CookieSecure:      false,
	CookieHTTPOnly:    false,
	CookieSameSite:    "",
	CookieSessionOnly: false,
	KeyGenerator: func() string {
		panic("TODO")
	},
	CookieName: "",
})

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

// create user session and store it in the context
func CreateUserSession(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Next()
	}
	sess, err := Store.Get(c)
	if err != nil {
		return err
	}
	sess.Set("user", user)
	if err := sess.Save(); err != nil {
		return err
	}
	return c.Next()
}

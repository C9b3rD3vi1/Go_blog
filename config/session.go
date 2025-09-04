package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/fiber/v2"
)

//var Store *session.Store

// InitSession initializes the session store with proper Fiber config
var Store = session.New(session.Config{
    Expiration:     time.Hour * 24, // Session expires after 24 hours
    KeyLookup:      "cookie:session_id",
    CookieSecure:   false,          // Set to true in production with HTTPS
    CookieHTTPOnly: true,           // JavaScript can't access the cookie
    CookieSameSite: "Lax",          // CSRF protection
    KeyGenerator:   utils.UUID,     // Proper key generator
    CookieName:     "session_id",   // Cookie name
})


// create user session and store it in the context
func CreateUserSession(c *fiber.Ctx) error {
	// fetch user from context
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

// Get current user from session
func GetCurrentUser(c *fiber.Ctx) *models.User {
	sess, err := Store.Get(c)
	if err != nil {
		return nil
	}
	user := sess.Get("user")
	if user == nil {
		return nil
	}
	return user.(*models.User)
}

package config

import (
	"fmt"
	"time"

	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

var Store *session.Store

// InitSession initializes the global session store
func InitSession() {
	Store = session.New(session.Config{
		Expiration:     24 * time.Hour, // Session expires after 24 hours
		KeyLookup:      "cookie:session_id",
		CookieSecure:   false,          // set true in production (HTTPS)
		CookieHTTPOnly: true,           // JS can't access cookie
		CookieSameSite: "Lax",          // CSRF protection
		KeyGenerator:   utils.UUID,     // unique session IDs
		CookieName:     "session_id",   // cookie name
	})
}

// CreateUserSession saves user into session
func CreateUserSession(c *fiber.Ctx) error {
    user, ok := c.Locals("user").(*models.User)
    if !ok || user == nil {
        fmt.Println("⚠️ No user in context, skipping session creation")
        return nil
    }

    sess, err := Store.Get(c)
    if err != nil {
        fmt.Println("❌ Error getting session:", err)
        return err
    }

    // store only user_id
    sess.Set("user_id", user.ID)
    if err := sess.Save(); err != nil {
        fmt.Println("❌ Error saving session:", err)
        return err
    }

    fmt.Printf("✅ Session created for user: ID=%d, Email=%s\n", user.ID, user.Email)
    fmt.Println("   Cookie:", sess.Keys())

    return nil
}


// GetCurrentUser fetches the logged-in user from session
func GetCurrentUser(c *fiber.Ctx) *models.User {
	sess, err := Store.Get(c)
	if err != nil {
		return nil
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return nil
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil
	}

	return &user
}

package config

import (
    "fmt"
    "time"

    "github.com/C9b3rD3vi1/Go_blog/database"
    "github.com/C9b3rD3vi1/Go_blog/models"
    "github.com/google/uuid"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/session"
    "github.com/gofiber/fiber/v2/utils"
)

var Store *session.Store

// InitSession initializes the global session store
func InitSession() {
    Store = session.New(session.Config{
        Expiration:     24 * time.Hour,
        KeyLookup:      "cookie:session_id",
        CookieSecure:   false,
        CookieHTTPOnly: true,
        CookieSameSite: "Lax",
        KeyGenerator:   utils.UUID,
    })
    fmt.Println("🟢 Session store initialized")
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

    sess.Set("user_id", user.ID.String())
    if err := sess.Save(); err != nil {
        fmt.Println("❌ Error saving session:", err)
        return err
    }

    fmt.Printf("✅ Session created for user: ID=%s, Email=%s\n", user.ID.String(), user.Email)
    return nil
}

// GetCurrentUser fetches the logged-in user from session
func GetCurrentUser(c *fiber.Ctx) *models.User {
    sess, err := Store.Get(c)
    if err != nil {
        fmt.Println("❌ Error fetching session:", err)
        return nil
    }

    idRaw := sess.Get("user_id")
    if idRaw == nil {
        fmt.Println("⚠️ No user_id found in session")
        return nil
    }

    idStr, ok := idRaw.(string)
    if !ok {
        fmt.Printf("⚠️ user_id is not a string: %v (%T)\n", idRaw, idRaw)
        return nil
    }

    userID, err := uuid.Parse(idStr)
    if err != nil {
        fmt.Println("❌ Invalid UUID in session user_id:", err)
        return nil
    }

    var user models.User
    if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
        fmt.Printf("❌ User not found in DB for ID=%s: %v\n", userID, err)
        return nil
    }

    fmt.Printf("✅ Current user fetched: ID=%s, Email=%s\n", user.ID, user.Email)
    return &user
}

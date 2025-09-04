package auth

import (
	"time"

	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

func AdminAuthHandler(c *fiber.Ctx) error {
    email := c.FormValue("email")
    password := c.FormValue("password")
    remember := c.FormValue("remember") == "on"

    // Find admin user
    var admin models.User
    if err := database.DB.Where("email = ? AND is_admin = ?", email, true).First(&admin).Error; err != nil {
        return c.Status(401).Render("admin/login", fiber.Map{
            "Error": "Invalid credentials",
        })
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
        return c.Status(401).Render("admin/login", fiber.Map{
            "Error": "Invalid credentials",
        })
    }

    // ✅ put user into context
    c.Locals("user", &admin)

    // ✅ save to session (via your helper)
    if err := config.CreateUserSession(c); err != nil {
        return c.Status(500).SendString("Error creating session")
    }

    // ✅ Handle "remember me" manually (optional)
    if remember {
        sess, err := config.Store.Get(c)
        if err == nil {
            sess.Set("remember", true)
            sess.SetExpiry(48 * time.Hour) // extend session
            _ = sess.Save()
        }
    }

    return c.Redirect("/admin/dashboard")
}



// Admin Logic OTP verification
func ShowOTPPage(c *fiber.Ctx) error {
	return c.Render("admin/verify", fiber.Map{})
}

// POST /admin/verify-otp
func VerifyOTPHandler(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("2fa_user_id")

	if userID == nil {
		return c.Redirect("/admin/login")
	}

	var user models.User
	err := database.DB.First(&user, userID).Error
	if err != nil {
		return c.Redirect("/admin/login")
	}

	otpCode := c.FormValue("otp")
	if !totp.Validate(otpCode, user.TwoFASecret) {
		return c.Render("admin/otp", fiber.Map{"Error": "Invalid OTP"})
	}

	// Set full admin session now
	sess.Delete("2fa_user_id")
	sess.Set("admin", user)
	sess.Save()

	return c.Redirect("/admin/dashboard")
}

// AdminLogoutHandler handles admin logout
func AdminLogoutHandler(c *fiber.Ctx) error {
	// Destroy the session
	sess, _ := config.Store.Get(c)
	sess.Destroy()

	return c.Redirect("/admin/login")
}


// AdminLoginForm handles admin login form
func AdminLoginForm(c *fiber.Ctx) error {
	return c.Render("admin/login", fiber.Map{})
}

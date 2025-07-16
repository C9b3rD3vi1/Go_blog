package auth

import (
	"time"

	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AdminAuthHandler(c *fiber.Ctx) error {
	// Get form values
	username := c.FormValue("username")
	password := c.FormValue("password")
	//otp := c.FormValue("otp") // For 2FA
	remember := c.FormValue("remember") == "on"

	// Find admin user
	var admin models.User

	if err := database.DB.Where("username = ? AND is_admin = ?", username, true).First(&admin).Error; err != nil {

		return c.Status(401).Render("admin/login", fiber.Map{
			"Error": "Invalid credentials",
		})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	// Set session
	sess, _ := config.Store.Get(c)
	sess.Set("2fa_user_id", admin.ID)

	// If remember is checked, set session expiry longer
	if remember {
		sess.SetExpiry(48 * time.Hour) // or any long session time
	}

	// save admin user to session
	if err := sess.Save(); err != nil {
		return c.Status(500).SendString("Error saving session")
	}

	return c.Redirect("/admin/verify")
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

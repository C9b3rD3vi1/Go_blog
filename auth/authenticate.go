package auth

import (
	//"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"

	//"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

// Define the routes
func UserRegisterHandler(c *fiber.Ctx) error {
	// Get form values
	Fullname := c.FormValue("fullname")
	Username := strings.TrimSpace(c.FormValue("username"))
	Email := strings.TrimSpace(c.FormValue("email"))
	Password := strings.TrimSpace(c.FormValue("password"))
	PasswordConfirm := strings.TrimSpace(c.FormValue("password_confirm"))

	// validate to make sure all fields are filled
	if Fullname == "" || Username == "" || Email == "" || Password == "" || PasswordConfirm == "" {
		return c.Render("pages/register", fiber.Map{
			"error": "All fields are required",
		})
	}

	// Check if passwords match
	if Password != PasswordConfirm {
		return c.Render("pages/register", fiber.Map{
			"error": "Passwords dont match",
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString("Error hashing password")
	}

	// Create a new user
	user := models.User{
		FullName: Fullname,
		Username: Username,
		Email:    Email,
		Password: string(hashedPassword),
	}

	// Save the user to the database
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).SendString("Error creating user")
	}
	// Redirect to the login page
	return c.Redirect("/login")
}

// UserLoginHandler handles user login
func UserLoginHandler(c *fiber.Ctx) error {
	// Get form values
	//username := strings.TrimSpace(c.FormValue("username"))
	email := strings.TrimSpace(c.FormValue("email"))
	password := strings.TrimSpace(c.FormValue("password"))

	// Find user in the database, using username
	var user models.User

	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return c.Render("/login", fiber.Map{
			"error": "Invalid email address!! Please try again",
		})
	}

	// Check hashedPassword password and compared to stored password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.Render("/pages/login", fiber.Map{
			"error": "Invalid username or password !! Please try again",
		})
	}

	// Create session
	sess, err := config.Store.Get(c)
	if err != nil {
		return err
	}
	// create and store user in session
	sess.Set("userID", user.ID)
	sess.Set("username", user.Username)
	sess.Set("_ip", c.IP())

	if err := sess.Save(); err != nil {
		return err
	}

	return c.Redirect("/")
}

// LogoutHandler handles user logout
func UserLogoutHandler(c *fiber.Ctx) error {
	sess, err := config.Store.Get(c)
	if err != nil {
		return err
	}
	sess.Destroy()
	return c.Redirect("pages/login")
}

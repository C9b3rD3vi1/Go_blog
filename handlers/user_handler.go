package handlers

import (
	//"fmt"
	"github.com/gofiber/fiber/v2"

	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/C9b3rD3vi1/Go_blog/config"
	"golang.org/x/crypto/bcrypt"
)


// Define the routes
func UserRegisterHandler(c *fiber.Ctx) error {
	// Get form values
	Fullname := c.FormValue("fullname")
	Username := c.FormValue("username")
	Email := c.FormValue("email")
	Password := c.FormValue("password")
	PasswordConfirm := c.FormValue("password_confirm")
	// Check if passwords match
	if Password != PasswordConfirm {
		return c.Status(400).SendString("Passwords do not match")
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
	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).SendString("Error creating user")
	}
	// Redirect to the login page
	return c.Redirect("/login")
}


// UserLoginHandler handles user login
func UserLoginHandler(c *fiber.Ctx) error {
	// Get form values
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	// Find the user in the database
	var user models.User
	result := config.DB.Where("username = ?", Username).First(&user)
	if result.Error != nil {
		return c.Status(401).SendString("Invalid username or password")
	}
	// Check if the password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if err != nil {
		return c.Status(401).SendString("Invalid username or password")
	}
	// Set the user session
	c.Locals("user", user)
	return c.Redirect("/")
}

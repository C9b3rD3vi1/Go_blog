package main

import (
	"fmt"
	"log"

	"github.com/C9b3rD3vi1/Go_blog/auth"
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/handlers"
	"github.com/C9b3rD3vi1/Go_blog/middleware"

	//"github.com/C9b3rD3vi1/Go_blog/models"

	//"github.com/C9b3rD3vi1/Go_blog/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/znbang/gofiber-layout/html"
)

// fibre app main function
func main() {
	// load template engine
	engine := html.New("./templates", ".html")

	//debug mode
	engine.Debug(true)

	// autoreload in dev environment
	engine.Reload(true)

	// Config app layouts
	engine.Layout("layouts/base")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//load static files
	app.Static("/static", "./static")

	// Initialize the database
	_, err := config.InitDB()
	if err != nil {
		log.Fatal("Could not initialize database:", err)
	}

	// admin login route
	app.Get("/admin/login", handlers.AdminAuthHandler)

	// handle post request to admin login
	app.Post("/admin/login", handlers.AdminAuthHandler)

	// Route to handle admin dashboard
	app.Get("/admin/dashboard", middleware.RequireAdminAuth, handlers.AdminDashboard, handlers.AdminCreatePost, handlers.AdminEditPostForm, handlers.AdminDeletePost)

	// Route to render index.html
	app.Get("/", handlers.HomePageHandler)

	//User registration route
	app.Get("/register", auth.UserRegisterHandler)
	app.Post("/register", auth.UserRegisterHandler)

	// Route to handle login
	app.Get("/login", auth.UserLoginHandler)
	// handle post request to login
	app.Post("/login", auth.UserLoginHandler)

	// Route to handle logout
	app.Get("/logout", auth.UserLogoutHandler)

	// app listen on port 3000
	fmt.Println("Server is running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/C9b3rD3vi1/Go_blog/auth"
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/handlers"
	"github.com/C9b3rD3vi1/Go_blog/middleware"

	//"github.com/C9b3rD3vi1/Go_blog/routes"
	//"github.com/C9b3rD3vi1/Go_blog/models"
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

	// time add function
	engine.AddFunc("now", func() string {
		return time.Now().Format("2006-01-02 15:04:05")
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//load static files
	app.Static("/static", "./static")
	app.Static("/upload", "./upload")

	// initialize session
	config.InitSession()

	// Initialize the database
	_, err := database.InitDB()
	if err != nil {
		log.Fatal("Could not initialize database:", err)
	}

	// admin login route
	app.Get("/admin/login", auth.AdminAuthHandler)

	// handle post request to admin login
	app.Post("/admin/login", auth.AdminAuthHandler)
	app.Get("/admin/verify", auth.ShowOTPPage)
	app.Post("/admin/verify", auth.ShowOTPPage)

	// Route to handle admin dashboard
	app.Get("/admin/dashboard", middleware.RequireAdminAuth, handlers.AdminDashboard,
		handlers.AdminEditPostForm, handlers.AdminDeletePost)

	// create blog post and save to database
	app.Get("/admin/create_blog", handlers.ShowCreateBlogForm)
	app.Post("/admin/create_blog", handlers.CreateBlogPostHandler)

	// Route to render index.html
	app.Get("/", handlers.HomePageHandler)

	//User registration route
	app.Get("/register", handlers.UserRegisterHandlerForm)
	app.Post("/register", auth.UserRegisterHandler)

	// Route to handle login
	app.Get("/login", handlers.UserLoginHandlerForm)
	// handle post request to login
	app.Post("/login", auth.UserLoginHandler)

	// contact
	app.Get("/contact", handlers.UserContactHandlerForm)
	//app.Post("/contact", handlers.UserContactHandler)

	// about us page
	app.Get("/about", handlers.AboutUsHandler)

	// Route to handle logout
	app.Get("/logout", auth.UserLogoutHandler)

	// Route to handle post request
	app.Post("/post", handlers.PostHandlerFunc)

	// blog
	app.Get("/blog", handlers.BlogHandler)
	app.Get("/blog_detail/:slug", handlers.BlogPostHandler)
	app.Get("/blog/:slug", handlers.BlogDetailsHandler)

	// github stats
	app.Get("/api/github-stats", handlers.GitHubStatsHandler)

	// app listen on port 3000
	fmt.Println("Server is running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

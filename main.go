package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/C9b3rD3vi1/Go_blog/auth"
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/handlers"
	"github.com/C9b3rD3vi1/Go_blog/routes"

	//"github.com/C9b3rD3vi1/Go_blog/middleware"

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
	app.Static("/uploads", "./uploads")

	// initialize session
	config.InitSession()

	// Initialize the database
 // Initialize database
    db, err := database.InitDB()
    if err != nil {
        log.Fatal("Database initialization failed:", err)
    }

    // Create admin user
    if err := database.CreateAdminUser(db); err != nil {
        log.Fatal("Failed to create admin user:", err)
    }
    log.Println("Admin user created/verified")


    
    
    // Setup Adminroutes
    routes.SetupAdminRoutes(app) 
    routes.SetupPublicRoutes(app)
    
    
//	app.Get("/admin/verify", auth.ShowOTPPage)
	//app.Post("/admin/verify", auth.ShowOTPPage)



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


	// github stats
	app.Get("/api/github-stats", handlers.GitHubStatsHandler)

	// app listen on port 3000
	fmt.Println("Server is running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

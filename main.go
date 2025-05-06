package main

import (
	"fmt"
    "log"

	"github.com/C9b3rD3vi1/Go_blog/handlers"
	"github.com/C9b3rD3vi1/Go_blog/models"
    "github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)


// fibre app main function
func main() {
    // load template engine
    engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{Views: engine})
    engine.Reload(true)

    fmt.Println("Server is running on port 3000")

      // Initialize the database
      _, err := config.InitDB()
      if err != nil {
          log.Fatal("Could not initialize database:", err)
      }


    // Route to render index.html
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
            "Title": "Hacker Hub!",})
    })


    // Route to handle posts
    app.Get("/posts", func(c *fiber.Ctx) error {
        post := models.CreateSamplePost()
        
        if (post == models.Post{}) {
            return c.Status(404).SendString("Post not found")
        }

        return c.Render("post", fiber.Map{
            "Post": post,
    })
    })

    //User registration route
    app.Get("/register", func(c *fiber.Ctx) error {
        return c.Render("register", fiber.Map{
            "Title": "Register",
        })
    })

    app.Post("/register", handlers.UserRegisterHandler)


    // Route to handle login
    app.Get("/login", func(c *fiber.Ctx) error {
        return c.Render("login", fiber.Map{
            "Title": "Login",
        })
    })

    // handle post request to login
    app.Post("/login", handlers.UserLoginHandler)

    
    // Route to handle logout
    app.Get("/logout", handlers.UserLogoutHandler)


    // app listen on port 3000
    app.Listen(":3000")
}
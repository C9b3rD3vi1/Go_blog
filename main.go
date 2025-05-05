package main

import (
	"fmt"
	//"net/http"

	"github.com/C9b3rD3vi1/Go_blog/models"
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
            "Title": post,
    })
    })

    // app listen on port 3000
    app.Listen(":3000")
}
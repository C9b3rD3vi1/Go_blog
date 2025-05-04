package main

import (
	"fmt"
	//"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"


)

// fibre app main function
func main() {
    // load template engine
    engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{Views: engine})

    fmt.Println("Server is running on port 3000")

    // Route to render index.html
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
            "Title": "Hello, World!, Welcome to the world of Go Fiber",})
    })

    // Route to handle posts
    app.Get("/posts", func(c *fiber.Ctx) error {
        return c.Render("posts,", fiber.Map{
            "Title": "Posts",
    })
    })


    // app listen on port 3000
    app.Listen(":3000")
}
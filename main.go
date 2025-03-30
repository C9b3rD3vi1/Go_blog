package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"


)


// fibre app main function
func main() {
    app := fiber.New()

    fmt.Println("Server is running on port 3000")

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Cyberlock Technologies!")
    })

    // app listen on port 3000
    app.Listen(":3000")
}
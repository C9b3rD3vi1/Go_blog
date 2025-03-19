package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"


)


// fibre app main function
func main() {
    app := fiber.New()

    fmt.Println("Server is running on port 3000")

    // configure routes
    // app listen on port 3000
    app.Listen(":3000")
}
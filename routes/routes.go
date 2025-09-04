package routes

import (
    "github.com/C9b3rD3vi1/Go_blog/handlers"
    "github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {
    // Public services pages
    app.Get("/services", handlers.ServiceList)        // List all services
    app.Get("/service/:id", handlers.ServiceView)    // Single service view

    // You can add other public pages here, e.g.:
    // app.Get("/", handlers.HomePage)
    // app.Get("/about", handlers.AboutPage)
}
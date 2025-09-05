package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/C9b3rD3vi1/Go_blog/models"
    
    "github.com/C9b3rD3vi1/Go_blog/database"
)

// Frontend handlers for services
func ServiceList(c *fiber.Ctx) error {
    var services []models.Services
    database.DB.Order("created_at desc").Find(&services)

    return c.Render("pages/services", fiber.Map{
        "Services": services,
        "Admin":    false, // public page, no admin controls
    })
}

// ServiceView displays a single service view
func ServiceView(c *fiber.Ctx) error {
    slug := c.Params("slug")
    var service models.Services
    if err := database.DB.First(&service, slug).Error; err != nil {
        return c.Status(404).SendString("Service not found")
    }

    return c.Render("pages/service_view", fiber.Map{
        "Service": service,
        "Admin":   false, // public view
    })
}
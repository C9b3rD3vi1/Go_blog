package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
)

// --- Services ---
// 
func AdminServiceList(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    var services []models.Services
    database.DB.Order("created_at desc").Find(&services)

    return c.Render("admin/services", fiber.Map{
        "Title":    "Services",
        "Admin":    admin.Username,
        "Services": services,
    })
}

// AdminCreateServices handles the creation of a new service
func AdminCreateServices(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    service := models.Services{
        Title:       c.FormValue("title"),
        Description: c.FormValue("description"),
        ImageURL:    c.FormValue("image"),
    }
    if err := database.DB.Create(&service).Error; err != nil {
        return c.Status(500).SendString("Error saving service")
    }

    return c.Redirect("/admin/dashboard")
}

// AdminEditServiceForm handles the form for editing a service
func AdminEditServiceForm(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    var service models.Services
    if err := database.DB.First(&service, id).Error; err != nil {
        return c.Status(404).SendString("Service not found")
    }

    return c.Render("admin/edit_service", fiber.Map{
        "Title":   "Edit Service",
        "Admin":   admin.Username,
        "Service": service,
    })
}

// AdminDeleteService handles the deletion of a service
func AdminDeleteService(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    if id == "" {
        return c.Status(404).SendString("Invalid Service ID")
    }

    if err := database.DB.Delete(&models.Services{}, id).Error; err != nil {
        return c.Status(500).SendString("Error deleting service")
    }

    return c.Redirect("/admin/dashboard")
}

// AdminUpdateService handles the update of a service
func AdminUpdateService(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    var service models.Services
    if err := database.DB.First(&service, id).Error; err != nil {
        return c.Status(404).SendString("Service not found")
    }

    service.Title = c.FormValue("title")
    service.Description = c.FormValue("description")
    service.ImageURL = c.FormValue("image")

    if err := database.DB.Save(&service).Error; err != nil {
        return c.Status(500).SendString("Error updating service")
    }

    return c.Redirect("/admin/services")
}

// AdminViewService handles the view of a service
func AdminViewService(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    var service models.Services
    if err := database.DB.First(&service, id).Error; err != nil {
        return c.Status(404).SendString("Service not found")
    }

    return c.Render("admin/view_service", fiber.Map{
        "Title":   service.Title,
        "Admin":   admin.Username,
        "Service": service,
    })
}


// AdminNewServiceForm renders the form to add a new service
func AdminNewServiceForm(c *fiber.Ctx) error {
    // Get current admin
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    return c.Render("admin/new_service", fiber.Map{
        "Title": "Add New Service",
        "Admin": admin.Username,
    })
}

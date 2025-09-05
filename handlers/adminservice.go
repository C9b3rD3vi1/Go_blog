package handlers

import (
	"strings"

	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/C9b3rD3vi1/Go_blog/utils"
	"github.com/gofiber/fiber/v2"
)

// --- Services ---
//// Use Preload("TechStacks") to eager load the many-to-many relationship
func AdminServiceList(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    var services []models.Services
    // Use Preload("TechStacks") to eager load the many-to-many relationship
    result := database.DB.Preload("TechStacks").Order("created_at desc").Find(&services)
    
    if result.Error != nil {
        // Handle the error appropriately, e.g., log it and return an error page.
        // For now, let's just show a simple error.
        return c.Status(500).SendString("Error loading services")
    }

    return c.Render("admin/services", fiber.Map{
        "Title":    "Services",
        "Admin":    admin.Username,
        "Services": services,
    })
}


func AdminNewServicePage(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    var techStacks []models.TechStack
    if err := database.DB.Order("created_at desc").Find(&techStacks).Error; err != nil {
        return c.Status(500).SendString("Error fetching tech stacks")
    }

    return c.Render("admin/new_service", fiber.Map{
        "Admin":      admin,
        "TechStacks": techStacks, // must match {{ .TechStacks }} in template
    })
}

// AdminCreateServices handles the creation of a new service
func AdminCreateServices(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    // Get form values
    title := c.FormValue("title")
    description := c.FormValue("description")


    // Get selected tech stack IDs from form (multiple checkboxes or select)
    stackIDs := c.FormValue("techstacks") // returns comma-separated if select[multiple]
    ids := strings.Split(stackIDs, ",")

    var techStacks []models.TechStack
    if len(ids) > 0 {
        database.DB.Where("id IN ?", ids).Find(&techStacks)
    }

    // Upload image if provided
    imageURL, _ := utils.UploadImage(c, "image")

    // Generate unique slug
    slug := utils.UniqueSlug(database.DB, "services", title)

    service := models.Services{
        Title:       title,
        Description: description,
        Slug:        slug,
        ImageURL:    imageURL,
        TechStacks:  techStacks,
    }

    if err := database.DB.Create(&service).Error; err != nil {
        return c.Status(500).SendString("Error saving service")
    }

    return c.Redirect("/admin/services") // 👈 better UX: go to services list
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

    slug := c.Params("slug")
    var service models.Services
    if err := database.DB.First(&service, slug).Error; err != nil {
        return c.Status(404).SendString("Service not found")
    }
    // Upload image if provided
    imageURL, _ := utils.UploadImage(c, "image")

    service.Title = c.FormValue("title")
    service.Description = c.FormValue("description")
    service.ImageURL = imageURL
    // Generate unique slug
    slug = utils.UniqueSlug(database.DB, "services", service.Title)

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

    slug := c.Params("slug")
    var service models.Services
    if err := database.DB.First(&service, slug).Error; err != nil {
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

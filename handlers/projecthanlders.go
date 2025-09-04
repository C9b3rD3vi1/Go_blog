package handlers

import (
    "github.com/C9b3rD3vi1/Go_blog/database"
    "github.com/C9b3rD3vi1/Go_blog/models"
    "github.com/gofiber/fiber/v2"
)

// List all projects (public)
func ProjectList(c *fiber.Ctx) error {
    var projects []models.Projects
    if err := database.DB.Order("created_at desc").Find(&projects).Error; err != nil {
        return c.Status(500).SendString("Error fetching projects")
    }

    return c.Render("pages/projects", fiber.Map{
        "Title":    "Projects",
        "Projects": projects,
    })
}

// View single project by slug (public)
func ProjectView(c *fiber.Ctx) error {
    slug := c.Params("slug")
    var project models.Projects

    if err := database.DB.Where("slug = ?", slug).First(&project).Error; err != nil {
        return c.Status(404).Render("errors/404", fiber.Map{"Message": "Project not found"})
    }

    return c.Render("pages/project_view", fiber.Map{
        "Title":   project.Title,
        "Project": project,
    })
}

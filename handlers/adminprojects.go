package handlers

import (
	"strings"

	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/C9b3rD3vi1/Go_blog/utils"
	"github.com/gofiber/fiber/v2"
)

// --- Projects ---

// Show new project form
func AdminNewProjectForm(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    return c.Render("admin/new_project", fiber.Map{
        "Title": "Add New Project",
        "Admin": admin,
    })
}


func AdminNewProjectPage(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    var techStacks []models.TechStack
    if err := database.DB.Order("created_at desc").Find(&techStacks).Error; err != nil {
        return c.Status(500).SendString("Error fetching tech stacks")
    }

    return c.Render("admin/new_project", fiber.Map{
        "Admin":      admin,
        "TechStacks": techStacks, // must match {{ .TechStacks }} in template
    })
}


// Handle project creation
func AdminCreateProject(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    title := c.FormValue("title")
    description := c.FormValue("description")
    link := c.FormValue("link")

    // Get selected tech stack IDs from form (multiple checkboxes or select)
    stackIDs := c.FormValue("techstacks") // returns comma-separated if select[multiple]
    ids := strings.Split(stackIDs, ",")

    var techStacks []models.TechStack
    if len(ids) > 0 {
        database.DB.Where("id IN ?", ids).Find(&techStacks)
    }

    // Upload image
    imageURL, _ := utils.UploadImage(c, "image")

    // Generate slug
    slug := utils.UniqueSlug(database.DB, "projects", title)

    project := models.Projects{
        Title:       title,
        Description: description,
        Link:        link,
        Slug:        slug,
        TechStacks:  techStacks,
        ImageURL:    imageURL,
    }

    if err := database.DB.Create(&project).Error; err != nil {
        return c.Status(500).SendString("Error saving project")
    }

    return c.Redirect("/admin/projects")
}

// List all projects
func AdminProjectList(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    var projects []models.Projects
    database.DB.Order("created_at desc").Find(&projects)

    return c.Render("admin/projects", fiber.Map{
        "Title":    "Manage Projects",
        "Admin":    admin,
        "Projects": projects,
    })
}


// View single project
func AdminViewProject(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    slug := c.Params("slug")
    var project models.Projects
    if err := database.DB.Where("slug = ?", slug).First(&project).Error; err != nil {
        return c.Status(404).Render("errors/404", fiber.Map{"Message": "Project not found"})
    }

    return c.Render("admin/view_project", fiber.Map{
        "Title":   "View Project",
        "Admin":   admin,
        "Project": project,
    })
}

// Show edit form
func AdminEditProjectForm(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    var project models.Projects
    if err := database.DB.First(&project, id).Error; err != nil {
        return c.Status(404).Render("errors/404", fiber.Map{"Message": "Project not found"})
    }

    return c.Render("admin/edit_project", fiber.Map{
        "Title":   "Edit Project",
        "Admin":   admin,
        "Project": project,
    })
}

// Handle project update
func AdminUpdateProject(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    var project models.Projects
    if err := database.DB.First(&project, id).Error; err != nil {
        return c.Status(404).Render("errors/404", fiber.Map{"Message": "Project not found"})
    }

    project.Title = c.FormValue("title")
    project.Description = c.FormValue("description")
    project.Link = c.FormValue("link")

    // Optional: update image if provided
    if imageURL, _ := utils.UploadImage(c, "image"); imageURL != "" {
        project.ImageURL = imageURL
    }

    if err := database.DB.Save(&project).Error; err != nil {
        return c.Status(500).SendString("Error updating project")
    }

    return c.Redirect("/admin/projects")
}

// Delete project
func AdminDeleteProject(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    if err := database.DB.Delete(&models.Projects{}, id).Error; err != nil {
        return c.Status(500).SendString("Error deleting project")
    }

    return c.Redirect("/admin/projects")
}

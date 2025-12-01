package handlers

import (
	"encoding/json"
	"time"
	"fmt"
	"mime/multipart"
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


// Handle project creation with enhanced struct
func AdminCreateProject(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    // Parse form values
    project := models.Projects{
        Title:            c.FormValue("title"),
        Description:      c.FormValue("description"),
        LongDescription:  c.FormValue("long_description"),
        ProblemStatement: c.FormValue("problem_statement"),
        SolutionApproach: c.FormValue("solution_approach"),
        KeyFeatures:      c.FormValue("key_features"),

        Link:       c.FormValue("link"),
        GithubLink: c.FormValue("github_link"),
        DemoLink:   c.FormValue("demo_link"),
        DocsLink:   c.FormValue("docs_link"),

        Category:    c.FormValue("category"),
        Difficulty:  c.FormValue("difficulty"),
        ProjectType: c.FormValue("project_type"),
        Tags:        c.FormValue("tags"),

        DevelopmentTime: c.FormValue("development_time"),
        TeamSize:       utils.ParseInt(c.FormValue("team_size")),
        LinesOfCode:    c.FormValue("lines_of_code"),
        Uptime:         c.FormValue("uptime"),
        ResponseTime:   c.FormValue("response_time"),
        UsersCount:     c.FormValue("users_count"),

        Featured:  c.FormValue("featured") == "on",
        Published: c.FormValue("published") == "on",
        Status:    c.FormValue("status"),
    }

    // Parse dates
    if completionDate := c.FormValue("completion_date"); completionDate != "" {
        if parsedDate, err := time.Parse("2006-01-02", completionDate); err == nil {
            project.CompletionDate = &parsedDate
        }
    }

    if startDate := c.FormValue("started_at"); startDate != "" {
        if parsedDate, err := time.Parse("2006-01-02", startDate); err == nil {
            project.StartedAt = &parsedDate
        }
    }

    // TechStacks handling
    stackIDs := c.FormValue("techstacks")
    if stackIDs != "" {
        ids := strings.Split(stackIDs, ",")
        var techStacks []models.TechStack
        database.DB.Where("id IN ?", ids).Find(&techStacks)
        project.TechStacks = techStacks
    }

    // Upload main image
    if imageURL, err := utils.UploadImage(c, "image"); err == nil && imageURL != "" {
        project.ImageURL = imageURL
    }

    // ---------------------------
    // FIXED: Upload multiple images from "gallery"
    // ---------------------------

    form, err := c.MultipartForm()
    if err == nil && form.File != nil {
        galleryFiles := form.File["gallery"]
        var galleryURLs []string
    
        for idx, file := range galleryFiles {
            tempField := fmt.Sprintf("gallery_%d", idx)
    
            // Insert file into the form under the temporary name
            form.File[tempField] = []*multipart.FileHeader{file}
    
            // Call UploadImage normally
            galleryURL, err := utils.UploadImage(c, tempField)
            if err == nil {
                galleryURLs = append(galleryURLs, galleryURL)
            }
    
            // Remove temp field
            delete(form.File, tempField)
        }
    
        if len(galleryURLs) > 0 {
            if galleryJSON, err := json.Marshal(galleryURLs); err == nil {
                project.Gallery = string(galleryJSON)
            }
        }
    }

    // Unique slug
    project.Slug = utils.UniqueSlug(database.DB, "projects", project.Title)

    // Save to DB
    if err := database.DB.Create(&project).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Error saving project: " + err.Error())
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
    // Use Preload("TechStacks") to eager load the many-to-many relationship
    result := database.DB.Preload("TechStacks").Order("created_at desc").Find(&projects)
    
    if result.Error != nil {
        // Handle the error appropriately, e.g., log it and return an error page.
        // For now, let's just show a simple error.
        return c.Status(500).SendString("Error loading projects")
    }

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

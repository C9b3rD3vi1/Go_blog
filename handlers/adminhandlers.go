package handlers

import (
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
)

// AdminDashboard renders the admin dashboard
func AdminDashboard(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil {
        return c.Redirect("/admin/login")
    }

    if !admin.IsAdmin {
        return c.SendStatus(fiber.StatusForbidden)
    }

    var posts []models.Post
    var projects []models.Projects
    var services []models.Services
    var users []models.User

    database.DB.Order("created_at desc").Find(&posts)
    database.DB.Order("created_at desc").Find(&projects)
    database.DB.Order("created_at desc").Find(&users)
    database.DB.Order("created_at desc").Find(&services)

    return c.Render("admin/dashboard", fiber.Map{
        "Title":    "Admin Dashboard",
        "Admin":    admin, // whole struct → you can display Username, Email, etc.
        "Posts":    posts,
        "Projects": projects,
        "Services": services,
        "Users":    users,
    })
}

// --- Posts ---
func AdminPostList(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    var posts []models.Post
    if err := database.DB.Find(&posts).Error; err != nil {
        return c.Status(500).SendString("Error fetching posts")
    }

    return c.Render("admin/posts", fiber.Map{
        "Title": "Admin Post List",
        "Admin": admin,
        "Posts": posts,
    })
}

func AdminEditPostForm(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    var post models.Post
    if err := database.DB.First(&post, id).Error; err != nil {
        return c.Status(404).Render("errors/404", fiber.Map{"Message": "Post not found"})
    }

    return c.Render("admin/edit", fiber.Map{
        "Post": post,
    })
}

func AdminUpdatePost(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    var post models.Post
    if err := database.DB.First(&post, id).Error; err != nil {
        return c.Status(404).Render("errors/404", fiber.Map{"Message": "Post not found"})
    }

    post.Title = c.FormValue("title")
    post.Content = c.FormValue("content")
    post.Slug = c.FormValue("slug")
    post.ImageURL = c.FormValue("image")
    post.Tags = c.FormValue("tags")
    post.Author = admin.Username

    if post.Title == "" || post.Slug == "" {
        return c.Render("admin/edit", fiber.Map{
            "Post":  post,
            "Error": "Title and Slug are required",
        })
    }

    if err := database.DB.Save(&post).Error; err != nil {
        return c.Status(500).SendString("Error updating post")
    }

    return c.Redirect("/admin/posts")
}

func AdminDeletePost(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    if err := database.DB.Delete(&models.Post{}, id).Error; err != nil {
        return c.Status(500).SendString("Error deleting post")
    }

    return c.Redirect("/admin/dashboard")
}

// --- Projects ---
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

func AdminCreateProject(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    project := models.Projects{
        Title:       c.FormValue("title"),
        Description: c.FormValue("description"),
        Link:        c.FormValue("link"),
        ImageURL:    c.FormValue("image"),
    }
    if err := database.DB.Create(&project).Error; err != nil {
        return c.Status(500).SendString("Error saving project")
    }

    return c.Redirect("/admin/dashboard")
}

func AdminDeleteProject(c *fiber.Ctx) error {
    admin := config.GetCurrentUser(c)
    if admin == nil || !admin.IsAdmin {
        return c.Redirect("/admin/login")
    }

    id := c.Params("id")
    if err := database.DB.Delete(&models.Projects{}, id).Error; err != nil {
        return c.Status(500).SendString("Error deleting project")
    }

    return c.Redirect("/admin/dashboard")
}

// --- Services ---
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

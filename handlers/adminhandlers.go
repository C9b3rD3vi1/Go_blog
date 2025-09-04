package handlers

import (
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/C9b3rD3vi1/Go_blog/utils"
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
    
    imageURL, _ := utils.UploadImage(c, "image")

    post.Title = c.FormValue("title")
    post.Content = c.FormValue("content")
    post.Slug = c.FormValue("slug")
    post.ImageURL = imageURL
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


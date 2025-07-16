package handlers

import (
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/middleware"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/session"
)

// AdminAuthHandler handles admin authentication

// AdminDashboard renders the admin dashboard
func AdminDashboard(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, _ := middleware.Store.Get(c)
	admin := sess.Get("admin")

	if admin == nil {
		return c.Redirect("/admin/login")
	}

	return c.Render("admin/dashboard", fiber.Map{
		"Title": "Admin Dashboard",
		"Admin": admin,
	})
}

// AdminPostList renders the admin post list
func AdminPostList(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, _ := middleware.Store.Get(c)
	admin := sess.Get("admin")

	if admin == nil {
		return c.Redirect("/admin/login")
	}

	// Fetch posts from the database
	var posts []models.Post
	result := database.DB.Find(&posts)
	if result.Error != nil {
		return c.Status(500).SendString("Error fetching posts")
	}

	return c.Render("admin/posts", fiber.Map{
		"Title": "Admin Post List",
		"Admin": admin,
		"Posts": posts,
	})
}

// AdminEditPostForm renders the admin post edit form
func AdminEditPostForm(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, _ := middleware.Store.Get(c)
	admin := sess.Get("admin")

	if admin == nil {
		return c.Redirect("/admin/login")
	}

	// Get post ID from URL parameters
	id := c.Params("id")

	// Fetch the post from the database
	var blogpost models.BlogPost
	if err := database.DB.First(&blogpost, id).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{
			"Message": "Post not found",
		})
	}
	return c.Render("admin/edit", fiber.Map{
		//"Post":    blogpost,
		//	"Message": "",
		//	"Error":   "",
	})
}

func AdminUpdatePost(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Redirect("/admin/login")
	}

	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}

	// Get post ID from URL parameters
	id := c.Params("id")

	// Fetch the post from the database
	var blogpost models.BlogPost
	if err := database.DB.First(&blogpost, id).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{
			"Message": "Post not found",
		})
	}

	// Get form values
	title := c.FormValue("title")
	content := c.FormValue("content")
	slug := c.FormValue("slug")
	image := c.FormValue("image")
	tags := c.FormValue("tags")

	// Optional: validate required fields
	if title == "" || slug == "" {
		return c.Render("admin/edit", fiber.Map{
			"Post":    blogpost,
			"Error":   "Title and Slug are required",
			"Message": "",
		})
	}

	// Update the post
	blogpost.Title = title
	blogpost.Content = content
	blogpost.Slug = slug
	blogpost.ImageURL = image // Corrected field name
	blogpost.Tags = tags
	blogpost.Author = admin.(models.User).Username

	// Save to DB
	if err := database.DB.Save(&blogpost).Error; err != nil {
		return c.Status(500).SendString("Error updating post")
	}

	return c.Redirect("/admin/posts")
}

// AdminDeletePost handles post deletion
func AdminDeletePost(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Redirect("/admin/login")
	}

	admin := sess.Get("admin")
	if admin == nil {
		return c.Redirect("/admin/login")
	}

	// Get post ID from URL parameters
	id := c.Params("id")

	// Fetch the post from the database
	var blogpost models.BlogPost
	if err := database.DB.First(&blogpost, id).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{
			"Message": "Post not found",
		})
	}

	// Delete the post from the database
	if err := database.DB.Delete(&blogpost).Error; err != nil {
		return c.Status(500).Render("errors/500", fiber.Map{
			"Error": "Error deleting post",
		})
	}

	return c.Redirect("/admin/dashboard") // Redirect to the posts page
}

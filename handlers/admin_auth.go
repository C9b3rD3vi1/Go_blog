package handlers

import (
	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/middleware"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	//"github.com/gofiber/fiber/v2/middleware/session"
)

// AdminAuthHandler handles admin authentication

func AdminAuthHandler(c *fiber.Ctx) error {
	// Get form values
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Find admin user
	var admin models.User

	result := config.DB.Where("username = ? AND is_admin = ?", username, true).First(&admin)
	if result.Error != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	// Set session
	sess, _ := middleware.Store.Get(c)
	sess.Set("admin", admin)
	sess.Save()

	return c.Redirect("/admin/dashboard")
}

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

// AdminLogoutHandler handles admin logout
func AdminLogoutHandler(c *fiber.Ctx) error {
	// Destroy the session
	sess, _ := middleware.Store.Get(c)
	sess.Destroy()

	return c.Redirect("/admin/login")
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
	result := config.DB.Find(&posts)
	if result.Error != nil {
		return c.Status(500).SendString("Error fetching posts")
	}

	return c.Render("admin/posts", fiber.Map{
		"Title": "Admin Post List",
		"Admin": admin,
		"Posts": posts,
	})
}

// AdminPostForm renders the admin post form
func AdminPostForm(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, _ := middleware.Store.Get(c)
	admin := sess.Get("admin")

	if admin == nil {
		return c.Redirect("/admin/login")
	}

	return c.Render("admin/post_form", fiber.Map{
		"Title": "Admin Post Form",
		"Admin": admin,
	})
}

// AdminCreatePost handles post creation
func AdminCreatePost(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, _ := middleware.Store.Get(c)
	admin := sess.Get("admin")

	if admin == nil {
		return c.Redirect("/admin/login")
	}

	// Get form values
	title := c.FormValue("title")
	content := c.FormValue("content")

	// Create a new post
	post := models.Post{
		Title:   title,
		Content: content,
		Slug:    title,        // You might want to generate a slug from the title
		Image:   "image.jpg",  // Handle image upload as needed
		Tags:    "tag1, tag2", // Handle tags as needed
		Author:  admin.(models.User).Username,
	}

	// Save the post to the database
	result := config.DB.Create(&post)
	if result.Error != nil {
		return c.Status(500).SendString("Error creating post")
	}

	return c.Redirect("/admin/posts")
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
	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).SendString("Post not found")
	}

	return c.Render("admin/post_edit", fiber.Map{
		"Title": "Admin Post Edit",
		"Admin": admin,
		"Post":  post,
	})
}

// AdminUpdatePost handles post update
func AdminUpdatePost(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, _ := middleware.Store.Get(c)
	admin := sess.Get("admin")

	if admin == nil {
		return c.Redirect("/admin/login")
	}

	// Get post ID from URL parameters
	id := c.Params("id")

	// Fetch the post from the database
	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).SendString("Post not found")
	}

	// Get form values
	title := c.FormValue("title")
	content := c.FormValue("content")
	slug := c.FormValue("slug")
	image := c.FormValue("image")
	tags := c.FormValue("tags")
	//categoryID := c.FormValue("category_id")

	// Update the post
	post.Title = title
	post.Content = content
	post.Slug = slug
	post.Image = image
	post.Tags = tags
	post.Author = admin.(models.User).Username

	// Save the updated post to the database
	result = config.DB.Save(&post)
	if result.Error != nil {
		return c.Status(500).SendString("Error updating post")
	}

	return c.Redirect("/admin/posts")
}

// AdminDeletePost handles post deletion
func AdminDeletePost(c *fiber.Ctx) error {
	// Get the admin user from the session
	sess, _ := middleware.Store.Get(c)
	admin := sess.Get("admin")

	if admin == nil {
		return c.Redirect("/admin/login")
	}

	// Get post ID from URL parameters
	id := c.Params("id")

	// Fetch the post from the database
	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).SendString("Post not found")
	}

	// Delete the post from the database
	result = config.DB.Delete(&post)
	if result.Error != nil {
		return c.Status(500).SendString("Error deleting post")
	}

	return c.Redirect("/admin/posts")
}

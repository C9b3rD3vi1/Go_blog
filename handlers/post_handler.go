package handlers

import (
	"strings"

	"github.com/C9b3rD3vi1/Go_blog/config"
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/C9b3rD3vi1/Go_blog/utils"
	"github.com/gofiber/fiber/v2"
)

// --- Posts ---

// List all posts
func AdminPostList(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	var posts []models.Post
	if err := database.DB.Order("created_at desc").Find(&posts).Error; err != nil {
		return c.Status(500).SendString("Error fetching posts")
	}

	return c.Render("admin/posts", fiber.Map{
		"Title": "Admin Post List",
		"Admin": admin,
		"Posts": posts,
	})
}

// Show new post form
func AdminNewPostForm(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	return c.Render("admin/new_post", fiber.Map{
		"Title": "Add New Post",
		"Admin": admin,
	})
}

// Handle post creation
func AdminCreatePost(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	title := c.FormValue("title")
	content := c.FormValue("content")
	tag := c.FormValue("tag")

	// Upload image
	imageURL, _ := utils.UploadImage(c, "image")

	// Generate slug
	slug := utils.UniqueSlug(database.DB, "posts", title)

	post := models.Post{
		Title:    title,
		Content:  content,
		Slug:     slug,
		ImageURL: imageURL,
		Tags:     []string{tag},
		Author: admin.Username,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		return c.Status(500).SendString("Error saving post")
	}

	return c.Redirect("/admin/posts")
}


// View single post
func AdminViewPosts(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	slug := c.Params("slug")
	var post models.Post
	if err := database.DB.Where("slug = ?", slug).First(&post).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{"Message": "Post not found"})
	}

	return c.Render("admin/view_post", fiber.Map{
		"Title": "View Post",
		"Admin": admin,
		"Post":  post,
	})
}

// Show edit form
func AdminEditPostsForm(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{"Message": "Post not found"})
	}

	return c.Render("admin/edit_post", fiber.Map{
		"Title": "Edit Post",
		"Admin": admin,
		"Post":  post,
	})
}

// Update post
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

	title := c.FormValue("title")
	content := c.FormValue("content")
	tags := strings.Split(c.FormValue("tags"), ",")

	// Update fields
	post.Title = title
	post.Content = content
	post.Tags = tags
	post.Slug = utils.UniqueSlug(database.DB, "posts", title)

	if imageURL, _ := utils.UploadImage(c, "image"); imageURL != "" {
		post.ImageURL = imageURL
	}

	// Validation
	if post.Title == "" || post.Slug == "" {
		return c.Render("admin/edit_post", fiber.Map{
			"Post":  post,
			"Error": "Title and Slug are required",
		})
	}

	if err := database.DB.Save(&post).Error; err != nil {
		return c.Status(500).SendString("Error updating post")
	}

	return c.Redirect("/admin/posts")
}

// Delete post
func AdminDeletePost(c *fiber.Ctx) error {
	admin := config.GetCurrentUser(c)
	if admin == nil || !admin.IsAdmin {
		return c.Redirect("/admin/login")
	}

	id := c.Params("id")
	if err := database.DB.Delete(&models.Post{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting post")
	}

	return c.Redirect("/admin/posts")
}

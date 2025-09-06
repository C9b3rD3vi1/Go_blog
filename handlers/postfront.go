package handlers

import (
	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
)

// Public: list all posts
func PublicPostList(c *fiber.Ctx) error {
	var posts []models.Post
	if err := database.DB.Order("created_at desc").Find(&posts).Error; err != nil {
		return c.Status(500).SendString("Error fetching posts")
	}


	return c.Render("pages/posts", fiber.Map{
		"Posts": posts,
	})
}



// Public: view single post
func PublicPostDetail(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var post models.Post
	if err := database.DB.Where("slug = ?", slug).First(&post).Error; err != nil {
		return c.Status(404).Render("errors/404", fiber.Map{
			"Message": "Post not found",
		})
	}

	return c.Render("pages/postdetail", fiber.Map{
		"Post": post,
	})
}

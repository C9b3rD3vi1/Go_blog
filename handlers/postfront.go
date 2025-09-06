package handlers

import (
	"bytes"
	"html/template"
	"log"

	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
	"github.com/yuin/goldmark"
)

// PublicPostList: list all posts with preloaded tags and generated excerpts
func PublicPostList(c *fiber.Ctx) error {
	var posts []models.Post

	// Fetch all posts with tags
	if err := database.DB.Preload("Tags").Order("created_at desc").Find(&posts).Error; err != nil {
		return c.Status(500).SendString("Error fetching posts")
	}

	// Generate excerpt for each post (first 150 chars of content)
	for i := range posts {
		content := posts[i].Content
		if len(content) > 150 {
			posts[i].Content = content[:150] + "..."
		}
	}

	return c.Render("pages/posts", fiber.Map{
		"Posts": posts,
	})
}



func PublicPostDetail(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var post models.Post

	// Log the incoming request slug
	log.Printf("Attempting to fetch post with slug: %s", slug)

	// Fetch the current post along with its tags
	if err := database.DB.Preload("Tags").Where("slug = ?", slug).First(&post).Error; err != nil {
		// Log the error if the post is not found
		log.Printf("Error fetching post '%s': %v", slug, err)
		return c.Status(404).Render("errors/404", fiber.Map{
			"Message": "Post not found",
		})
	}

	// Log a success message for the fetched post
	log.Printf("Successfully fetched post: '%s' (ID: %d)", post.Title, post.ID)

	// Convert Markdown to HTML
	md := goldmark.New()
	var buf bytes.Buffer
	if err := md.Convert([]byte(post.Content), &buf); err != nil {
		// Log the error if Markdown conversion fails
		log.Printf("Error converting Markdown for post '%s': %v", post.Title, err)
		return c.Status(500).SendString("Error rendering post content")
	}

	// Fetch related posts: at least one tag in common
	var relatedPosts []models.Post
	if len(post.Tags) > 0 {
		// Log that the function will now search for related posts
		log.Printf("Searching for related posts for '%s' with %d tags...", post.Title, len(post.Tags))

		// Collect tag IDs
		tagIDs := make([]uint, len(post.Tags))
		for i, t := range post.Tags {
			tagIDs[i] = t.ID
		}

		// Join post_tags and select posts with at least one matching tag
		if err := database.DB.
			Joins("JOIN post_tags ON post_tags.post_id = posts.id").
			Where("post_tags.tag_id IN ?", tagIDs).
			Where("posts.id != ?", post.ID).
			Preload("Tags").
			Group("posts.id").
			Limit(5).
			Find(&relatedPosts).Error; err != nil {
			log.Printf("Error fetching related posts for '%s': %v", post.Title, err)
			relatedPosts = []models.Post{}
		}

		// Log the number of related posts found
		log.Printf("Found %d related posts for '%s'.", len(relatedPosts), post.Title)
	} else {
		log.Printf("Post '%s' has no tags, skipping related posts search.", post.Title)
	}

	return c.Render("pages/postdetail", fiber.Map{
		"Post":         post,
		"ContentHTML":  template.HTML(buf.String()),
		"RelatedPosts": relatedPosts,
	})
}

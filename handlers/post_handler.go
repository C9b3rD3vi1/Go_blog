package handlers

import (
	"encoding/json"

	//"fmt"
	"net/http"

	"github.com/C9b3rD3vi1/Go_blog/database"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"github.com/gofiber/fiber/v2"
)

// Define the routes
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the request
	var post models.Post
	//post := models.CreateSamplePost()

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set category base on Category
	//post.Category = post.CategoryID

	// save the post to the database
	result := database.DB.Create(&post)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)

}

// ShowPostHandler handles the request to show a post
func ShowPostHandler(c *fiber.Ctx) error {
	post := models.CreateSamplePost() // or fetch from DB

	return c.Render("post", post)
}

func PostHandlerFunc(c *fiber.Ctx) error {
	//	post := models.CreateSamplePost() // or fetch from DB

	return c.Render("pages/post", fiber.Map{
		//"post": post,
	})
}

func GitHubStatsHandler(c *fiber.Ctx) error {
	resp, err := http.Get("https://api.github.com/repos/C9b3rD3vi1/Go_blog/contributors")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "GitHub API error"})
	}
	defer resp.Body.Close()

	var contributors []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&contributors)

	return c.JSON(fiber.Map{
		"contributors": len(contributors),
	})
}

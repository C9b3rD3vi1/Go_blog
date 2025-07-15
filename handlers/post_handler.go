package handlers

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"github.com/C9b3rD3vi1/Go_blog/config"
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
	result := config.DB.Create(&post)

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

func HomePageHandler(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{})

}

func UserRegisterHandlerForm(c *fiber.Ctx) error {
	return c.Render("pages/register", fiber.Map{})

}

func UserLoginHandlerForm(c *fiber.Ctx) error {
	return c.Render("pages/login", fiber.Map{})

}

func UserContactHandlerForm(c *fiber.Ctx) error {
	return c.Render("pages/contact", fiber.Map{})

}

func AboutUsHandler(c *fiber.Ctx) error {
	return c.Render("pages/about", fiber.Map{})

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

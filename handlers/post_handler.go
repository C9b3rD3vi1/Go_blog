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

func HomePageHandler(c *fiber.Ctx) error  {
	return c.Render("/", fiber.Map{})

}

func UserRegisterHandler)(c *fiber.Ctx) error  {
	return c.Render("pages/register", fiber.Map{},)

}

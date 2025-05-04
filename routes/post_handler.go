package routes

import (
	"encoding/json"
	"fmt"
	"logger"
	"net/http"

	"github.com/C9b3rD3vi1/Go_blog/models"
)

// Define the routes
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the request
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set category base on Category
	//post.Category = post.CategoryID

	// 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)

	// save the post to the database
	//db.Create(&post)

	// Check for errors
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	// Log the post creation
	logger.Println("Post created:", post)

}
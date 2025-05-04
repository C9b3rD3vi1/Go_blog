package models

import (
	"time"

	"golang.org/x/text/internal/tag"
)

// Post represents a blog post in the database
type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Slug      string    `json:"slug" gorm:"not null"`
	Image     string    `json:"image" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Tags      string    `json:"tags" gorm:"not null"`
	Category  Category  `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}



// category represents a blog post category in the database
type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}


// CreatePost creates a new post in the database
func CreatePost(ID uint, title string, slug string, image string, content string, tags string) Post {
	post := Post{
		Title:    title,
		Slug:     slug,
		Image:    image,
		Content:  content,
		Tags:     tags,
		Category: Category{
			Name: "Category 1",
		},
	}
	return post
}
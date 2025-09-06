package models

import (
	"time"
	//"golang.org/x/text/internal/tag"
)

// Post represents a blog post in the database
// Post represents a blog post in the database
type Post struct {
    ID        uint      `gorm:"primaryKey"`
    Title     string    `gorm:"not null"`
    Slug      string    `gorm:"not null;uniqueIndex"`
    ImageURL  string    `gorm:"not null"`
    Content   string    `gorm:"not null"`
    Author    string    `gorm:"not null"`
    Tags      []Tag     `gorm:"many2many:post_tags;"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Tag represents a tag in the database
type Tag struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null;uniqueIndex"`
    Posts     []Post    `gorm:"many2many:post_tags;"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

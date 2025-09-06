package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	//"golang.org/x/text/internal/tag"
)

// Post represents a blog post in the database
type Post struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"not null"`
	Slug       string `gorm:"not null;uniqueIndex"`
	ImageURL     string `gorm:"not null"`
	Content    string `gorm:"not null"`
	Author     string `gorm:"not null"`
 	Tags      datatypes.JSONSlice[string] `gorm:"type:json"` // cross-db support


	// CategoryID is the foreign key for the category
	//Category  []Category  `gorm:"foreignKey:CategoryID"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	gorm.Model
}


// category represents a blog post category in the database
type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

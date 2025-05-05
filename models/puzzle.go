package models

import (
	"time"
	//"gorm.io/gorm"
	//"gorm.io/gorm/logger"

)

//  Create  a puzzle struct
type Puzzle struct {
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


// PuzzleCategory 
type PuzzleCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}


// CreatePuzzle creates a new puzzle in the database
func CreatePuzzle(ID uint, title string, slug string, image string, content string, tags string, category Category) Puzzle {
	puzzle := Puzzle{
		Title:    title,
		Slug:     slug,
		Image:    image,
		Content:  content,
		Tags:     tags,
		Category: category,
	}

	return puzzle
}
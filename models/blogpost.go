package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPost struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Slug      string    `json:"slug" gorm:"not null;uniqueIndex"`
	Excerpt   string    `json:"excerpt" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Tags      string    `json:"tags" gorm:"not null;default:''"`
	Published time.Time `json:"published" gorm:"not null"`
	Content   string
}

var now = time.Now()

func (b *BlogPost) CreateBlog(Title string, Slug string, Excerpt string, ImageURL string, Author string, Published string) BlogPost {
	blogpost := BlogPost{
		ID:        uuid.New(),
		Title:     "Sample Post",
		Slug:      "sample-post",
		Excerpt:   "This is a sample post content.",
		ImageURL:  "https://example.com/sample.jpg",
		Author:    "John Doe",
		Published: now,
	}
	return blogpost

}

func (b *BlogPost) BeforeCreate(tx *gorm.DB) (err error) {
    if b.ID == uuid.Nil {
        b.ID = uuid.New()
    }
    return
}
package models

import (
	"time"
)

type BlogPost struct {
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
		Title:     "Sample Post",
		Slug:      "sample-post",
		Excerpt:   "This is a sample post content.",
		ImageURL:  "https://example.com/sample.jpg",
		Author:    "John Doe",
		Published: now,
	}
	return blogpost

}

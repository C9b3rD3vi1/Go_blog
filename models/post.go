package models

import (
	"strings"
	"time"
	"unicode"
	"strconv"

	"gorm.io/gorm"
	//"golang.org/x/text/internal/tag"
)

// Post represents a blog post in the database
type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Slug      string    `json:"slug" gorm:"not null;uniqueIndex"`
	Image     string    `json:"image" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Tags      string    `json:"tags" gorm:"not null"`
	CategoryID uint      `json:"category_id" gorm:"not null"`
	// CategoryID is the foreign key for the category
	Category  Category  `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	gorm.Model
}



// category represents a blog post category in the database
type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}


// CreatePost creates a new post in the database
func CreatePost(ID uint, title string, slug string, image string, content string, tags string, category Category) Post {
	post := Post{
		Title:    title,
		Slug:     slug,
		Image:    image,
		Content:  content,
		Tags:     tags,
		Category: category,
	}

	return post
}


// Create category creates a new category in the database
func CreateCategory(ID uint, name string) Category {
	category := Category{
		Name: name,
	}

	return category
}




// create a sample post
func CreateSamplePost() Post {
	post := Post{
		Title:    "Sample Post",
		Slug:     "sample-post",
		Image:    "https://example.com/sample.jpg",
		Content:  "This is a sample post content.",
		Tags:     "sample, post",
		Category: Category{Name: "Sample Category"},
	}

	return post
}

// GenerateSlug generates a slug from the title
func GenerateSlug(title string) string {
	var sb strings.Builder
	// LOwercase , Rmove non-letters, and replace spaces with hyphens
	for _, r := range title {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			sb.WriteRune(r)
		} 
	}
	return strings.ToLower(strings.ReplaceAll(sb.String(), " ", "_")) // Convert to lowercase
}

// Hook GenerateSlug to the Post model
func (p *Post) BeforeSave(tx *gorm.DB) (err error) {
	if p.Slug == ""{

	 p.Slug = GenerateSlug(p.Title)

	}
	return
}


// Hook GenerateSlug to the Post model
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	baseSlug := GenerateSlug(p.Title)

	Slug := baseSlug
	
	var count int64
	i := 1

	// Loop until a unique slug is found
	for {
	// Check if the slug already exists in the database
	tx.Model(&Post{}).Where("slug = ?", Slug).Count(&count)
	if count == 0 {
		break
	}
	Slug = baseSlug + "_" + strconv.Itoa(i)
	i++
	}
	p.Slug = Slug

	return nil
}
package models

import (
	"fmt"

	"gorm.io/gorm"
)

// User struct represents a user entity with personal and contact details.
type User struct {
	FullName        string `gorm:"unique;not null"`
	ID              int    `gorm:"primaryKey"`
	Username        string `gorm:"unique;not null"`
	Email           string `gorm:"unique;not null"`
	Password        string `gorm:"required"`
	PasswordConfirm string
	Address         string
	TwoFASecret     string // save TOTP secret here
	IsActive        bool

	// admin
	IsAdmin bool `gorm:"default:false"`

	gorm.Model
}

// Comment struct represents a comment entity with user and post details.
type Comment struct {
	ID      int `gorm:"primaryKey"`
	Content string
	// UserID is the foreign key for the user
	UserID int
	User   User
	// PostID is the foreign key for the post
	PostID int
	Post   Post

	gorm.Model
}

// Create User function creates a new user entity.
func CreateUser(ID int, fullName string, username string, email string, address string, isActive bool) User {
	user := User{
		FullName: fullName,
		ID:       ID,
		Username: username,
		Email:    email,
		Address:  address,
		IsActive: isActive,
	}

	return user
}

// UserCreate function creates a new user entity.
func UserCreate() {
	user := User{
		FullName: "Nana Kwame",
		ID:       1,
		Email:    "nana.kwame@gmail.com",
		Username: "nana_kwame",
		Address:  "123 Main St, Anytown, USA",
		IsActive: true,
	}

	fmt.Println(user)
	fmt.Println("User created")
}

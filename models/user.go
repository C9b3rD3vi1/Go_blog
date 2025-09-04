package models

import (
	"golang.org/x/crypto/bcrypt"
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


// HashPassword hashes the user's password
func (u *User) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

// CheckPassword compares plain password with hashed password
func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}

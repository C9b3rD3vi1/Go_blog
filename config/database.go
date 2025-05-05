package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"github.com/C9b3rD3vi1/Go_blog/models"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB
// DBConnection is the function to connect to the database

func InitDB() (*gorm.DB, error) {
	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open("server.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		//return nil, err
	}
	// Set the global DB variable
	DB = db
	// Log the database connection
	log.Println("Connected to the database")
	
	// Migrate the schema
	if err := db.AutoMigrate(&models.Post{}, &models.Category{}, &models.User{}, &models.Comment{}); err != nil {
		log.Fatal("Failed to migrate the database schema:", err)
		return nil, err
	}
	// Log the migration
	log.Println("Database schema migrated successfully")
	
	return db, nil
}

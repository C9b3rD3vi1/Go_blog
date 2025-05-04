package config

import (
	"github.com/C9b3rD3vi1/Go_blog/models"
	"gorm.io/gorm"
	//"fmt"
	//"log"
)

// DB is the database connection
var DB *gorm.DB
// DBConnection is the function to connect to the database
func InitDB() (*gorm.DB, error) {
	// Connect to the SQLite database
	db, err := gorm.Open()
	if err != nil {
		return nil, err
	}
	
	// Migrate the schema
	db.AutoMigrate(&models.Post{}, &models.Category{}, &models.User{}, &models.Comment{})
	
	return db, nil
}




// ConnectDB connects to the SQLite database
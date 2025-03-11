package config

import (
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"fmt"
	//"log"
)

// DB is the database connection
var DB *gorm.DB

// ConnectDB connects to the SQLite database
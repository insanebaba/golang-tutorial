package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// DB to be used across application for database
	DB *gorm.DB
)

// Initialize database connection for global usage
func Initialize() {
	var err error
	DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to  connect to Database books.db")
	}
	DB.AutoMigrate(&Book{})

	for _, book := range Books {
		DB.Create(&book)
	}

}

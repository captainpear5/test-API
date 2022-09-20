package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var database *gorm.DB

func connectDatabase() {
	sqlDatabase, err := gorm.Open("sqlite3", "post.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	sqlDatabase.AutoMigrate(&Post{})

	database = sqlDatabase
}

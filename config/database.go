package config

import (
	"github.com/mhdianrush/go-fiber-books-rest-api/entities"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var logger = logrus.New()

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:admin@tcp(127.0.0.1:3306)/go_fiber_books_rest_api"), &gorm.Config{})
	if err != nil {
		logger.Println("Failed to Connect Database")
	}
	db.AutoMigrate(&entities.Book{})
	
	logger.Println("Database Connected")
	DB = db
}

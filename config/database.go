package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-fiber-books-rest-api/entities"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var logger = logrus.New()

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		logger.Printf("failed connect to database %s", err.Error())
	}
	db.AutoMigrate(&entities.Book{})

	logger.Println("database connected")
	DB = db
}

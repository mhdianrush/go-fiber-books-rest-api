package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-fiber-books-rest-api/config"
	"github.com/mhdianrush/go-fiber-books-rest-api/controllers"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", controllers.Index)
	book.Get("/:id", controllers.Find)
	book.Post("/", controllers.Create)
	book.Put("/:id", controllers.Update)
	book.Delete("/:id", controllers.Delete)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Println(err.Error())
	}
	logger.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	if err = app.Listen(":" + os.Getenv("SERVER_PORT")); err != nil {
		logger.Printf("failed connect to server %s", err.Error())
	}
	
	logger.Printf("Server Running on Port %s", os.Getenv("SERVER_PORT"))
}

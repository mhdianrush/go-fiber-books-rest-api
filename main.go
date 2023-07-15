package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
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

	logger.Println("Server Running on Port 8080")

	app.Listen(":8080")
}

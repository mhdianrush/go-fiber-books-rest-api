package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhdianrush/go-fiber-books-rest-api/config"
	"github.com/mhdianrush/go-fiber-books-rest-api/entities"
)

func Index(c *fiber.Ctx) error {
	var books []entities.Book

	config.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(books)
}

func Find(c *fiber.Ctx) error {
	return nil
}

func Create(c *fiber.Ctx) error {
	var book entities.Book
	
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := config.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhdianrush/go-fiber-books-rest-api/config"
	"github.com/mhdianrush/go-fiber-books-rest-api/entities"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var books []entities.Book

	if err := config.DB.Find(&books).Error; err != nil {
		// have to use return
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func Find(c *fiber.Ctx) error {
	id := c.Params("id")
	var book entities.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Not Found Book Data",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(book)
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
	id := c.Params("id")
	var book entities.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if config.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "can't update the data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully update book's data",
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var book entities.Book

	if config.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "can't delete book's data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully deleted book's data",
	})
}

package router

import (
	"github.com/api-auth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func BooksRouter(route fiber.Router, db *gorm.DB) {

	// get-all-books
	route.Get("/", func(c *fiber.Ctx) error {

		var books []models.Book

		if err := db.Find(&books).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"message": "server failed",
					"error":   err.Error(),
				})
		}

		return c.Status(fiber.StatusOK).
			JSON(books)
	})

	// get-book-by-id
	route.Get("/:id", func(c *fiber.Ctx) error {
		bookId, _ := c.ParamsInt("id")

		book := new(models.Book)

		if err := db.Where("id = ?", bookId).First(&book).Error; err != nil {
			return c.Status(fiber.StatusNotFound).
				JSON("Book your Found")
		}

		return c.Status(fiber.StatusOK).JSON(book)
	})

	// post-book
	route.Post("/", func(c *fiber.Ctx) error {
		book := new(models.Book)
		book.UserID = int(c.Locals("userId").(float64))

		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		}

		if err := db.Create(&book).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"errro": err.Error(),
				})
		}

		return c.Status(fiber.StatusCreated).
			JSON(fiber.Map{
				"status": "success",
				"book":   book,
			})
	})

	// post-book-by-id
	route.Patch("/:id", func(c *fiber.Ctx) error {
		userId := c.Locals("userId")
		bookId, _ := c.ParamsInt("id")
		var book models.Book
		if err := db.Where("id = ? AND user_id = ?", bookId, userId).First(&book).Error; err != nil {
			return c.Status(fiber.StatusConflict).
				JSON(fiber.Map{
					"error": "Book not found",
				})
		}

		var input models.UpdateBookRequest

		if err := c.BodyParser(&input); err != nil {
			return fiber.ErrBadRequest
		}

		if input.Title != nil {
			book.Title = *input.Title
		}

		if input.Description != nil {
			book.Description = *input.Description
		}

		if input.Price != nil {
			book.Price = *input.Price
		}

		if err := db.Save(&book).Error; err != nil {
			return fiber.ErrInternalServerError
		}

		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "Book Updated",
				"book":    book,
			})
	})

	// delete-book-by-id
	route.Delete("/:id", func(c *fiber.Ctx) error {
		bookId, _ := c.ParamsInt("id")
		userId := c.Locals("userId")

		log.Infof("DELETE bookId=%d userId=%d", bookId, userId)

		var book models.Book

		if err := db.Where("id = ? AND user_id = ?", bookId, userId).First(&book).Error; err != nil {
			return c.Status(fiber.StatusNotFound).
				JSON(fiber.Map{
					"error": "Book not found",
				})
		}

		if err := db.Delete(&book).Error; err != nil {
			return fiber.ErrInternalServerError
		}

		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "book deleted sucessfull",
			})
	})

}

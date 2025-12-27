package router

import (
	"github.com/api-auth/models"
	"github.com/gofiber/fiber/v2"
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
			JSON(fiber.Map{
				"message": "get all books",
			})
	})

	// get-book-by-id
	route.Get("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
			})
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
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
			})
	})

	// delete-book-by-id
	route.Delete("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
			})
	})

}

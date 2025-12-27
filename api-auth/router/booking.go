package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BooksRouter(route fiber.Router, db *gorm.DB) {

	// get-all-books
	route.Get("/", func(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
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

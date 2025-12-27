package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BooksRouter(route fiber.Router, db *gorm.DB) {
	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
			})
	})

	route.Get("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
			})
	})

	route.Post("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
			})
	})

	route.Patch("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
			})
	})

	route.Delete("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "get all books",
			})
	})

}

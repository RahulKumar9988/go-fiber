package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("2 leaning about routing")
	})

	// Grouping Routes
	app.Route("/v1/users", func(api fiber.Router) {
		api.Get("/", handler)
		api.Post("/", handler)
		api.Patch("/:id", handler)
		api.Delete("/:id", handler)
	})

	app.Server().MaxConnsPerIP = 1
	app.Listen(":3000")
}

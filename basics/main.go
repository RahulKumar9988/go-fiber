package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// simple server
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome to fiber server")
	})

	// params
	app.Get("/v1/:input", func(c *fiber.Ctx) error {
		return c.SendString("Input Value is : " + c.Params("input"))
	})
	app.Listen(":3000")
}

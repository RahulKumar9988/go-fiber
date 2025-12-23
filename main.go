package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	//simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to go fiber Server")
	})

	//params
	app.Get("/input/:value", func(c *fiber.Ctx) error {
		return c.SendString("Entred value is : " + c.Params("value"))
	})

	//query-paramter
	app.Get("/name/:value?", func(c *fiber.Ctx) error {
		if c.Params("value") != "" {
			return c.SendString("Hello! " + c.Params("value"))
		} else {
			return c.SendString("where is RJ.")
		}
	})

	// dynamic-route
	app.Get("/path/*", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("*"))
	})

	app.Listen(":3000")
}

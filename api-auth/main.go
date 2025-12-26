package main

import (
	"github.com/api-auth/auth"
	"github.com/api-auth/db"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// initialization of db
	db := db.InitializeDB()

	// creating fiber app instance
	app := fiber.New(fiber.Config{
		AppName: "Library API",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to library api backend")
	})

	// auth-handler
	auth.AuthHandlers(app.Group("/auth"), db)

	// listen app on port 3000
	app.Listen(":3000")
}

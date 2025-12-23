package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name" db:"namae"`
	Age  int    `json:"age" db:"age"`
}

type ErrorResponse struct {
	Message string
	Status  string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("leaning about routing")
	})

	app.Get("/store/:storeName/:product", func(c *fiber.Ctx) error {
		storeName := c.Params("storeName")
		product := c.Params("product")

		fmt.Println(storeName, " + ", product)

		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/pars", func(c *fiber.Ctx) error {
		p := new(Person)
		if err := c.BodyParser(p); err != nil {
			c.Locals("message : ", "Error found", err)
			return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
				Message: "Pass the body",
				Status:  "400",
			})
		}

		return c.Status(fiber.StatusOK).JSON(Person{
			Name: p.Name,
			Age:  p.Age,
		})

	})
	app.Server().MaxConnsPerIP = 2
	app.Listen(":3000")
}

package main

import (
	"fmt"
	"time"

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

//cookies struct

type Cookie struct {
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	Path        string    `json:"path"`
	Domain      string    `json:"domain"`
	MaxAge      int       `json:"max_age"`
	Expires     time.Time `json:"expires"`
	Secure      bool      `json:"secure"`
	HTTPOnly    bool      `json:"http_only"`
	SameSite    string    `json:"same_site"`
	SessionOnly bool      `json:"session_only"`
}

func main() {
	app := fiber.New()

	app.Get("/login", func(c *fiber.Ctx) error {

		cookieAuth := new(fiber.Cookie)
		cookieAuth.Name = "rahul"
		cookieAuth.Value = "asda@#@$ASD%11"
		cookieAuth.MaxAge = 100
		cookieAuth.Expires = time.Now().Add(24 * time.Hour)
		c.Cookie(cookieAuth)
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/checkout", func(c *fiber.Ctx) error {
		c.Cookies("rahul")
		fmt.Println("rahul : ", c.Cookies("rahul"))
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		c.ClearCookie()
		return c.SendStatus(fiber.StatusOK)
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

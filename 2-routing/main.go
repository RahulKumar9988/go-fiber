package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Verification struct {
	Name    string `json:"name" db:"name"`
	Role    string `json:"role" db:"role"`
	IsHuman bool   `json:"isHuman" db:"isHuman"`
}

type ErrorResponse struct {
	Message string
	Status  string
}

// cookies struct
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

type QueryUser struct {
	ID      int    `query:"id"`
	Dpt     string `query:"dpt"`
	IsAdmin bool   `query:"isAdmin"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		user := new(QueryUser)

		if err := c.QueryParser(user); err != nil {
			return err
		}

		log.Println("id", user.ID)
		log.Println("dpt", user.Dpt)
		log.Println("isAdmin", user.IsAdmin)

		return c.Status(fiber.StatusOK).JSON(user)
	})

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
	app.Server().MaxConnsPerIP = 2
	app.Listen(":3000")
}

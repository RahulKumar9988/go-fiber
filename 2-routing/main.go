package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
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

func rndNumber() int {
	id := rand.Int()
	fmt.Printf("id is : ", id)
	return id
}

func main() {
	app := fiber.New()

	// playing with files

	app.Post("/upload", func(c *fiber.Ctx) error {

		file, err := c.FormFile("profile")
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).
				JSON(fiber.Map{
					"message": "file is required",
				})
		}

		if file.Size > 2*1024*1024 {
			return c.Status(fiber.ErrBadRequest.Code).
				JSON(fiber.Map{
					"message": "size should be less then 2MB",
				})
		}

		allowType := map[string]bool{
			"image/webp": true,
			"image/png":  true,
		}

		if !allowType[file.Header.Get("Content-Type")] {
			return c.Status(fiber.ErrBadRequest.Code).
				JSON(fiber.Map{
					"message": "filetype not matched",
				})
		}

		uplaodDir := "./uploads"
		os.MkdirAll(uplaodDir, os.ModePerm)

		ext := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		path := filepath.Join(uplaodDir, fileName)

		if err := c.SaveFile(file, path); err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"message": "failed to upload file",
				})
		}

		return c.JSON(fiber.Map{
			"message": "file uplaod sucessfully",
			"file":    fileName,
		})
	})

	app.Post("/bulk-upload", func(c *fiber.Ctx) error {

		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": "failed to pars file",
				})
		}

		files := form.File["notes"]

		if len(files) == 0 {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": "at list 1 file required",
				})
		}

		fileType := map[string]bool{
			"image/webp": true,
			"image/png":  true,
		}

		for _, file := range files {
			if !fileType[file.Header.Get("Content-Type")] {
				return c.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"message": "file doesnot supported",
					})
			}

		}

		uplaodDir := "./uploads"
		if err := os.MkdirAll(uplaodDir, os.ModePerm); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}

		var uploadFilesLists []string

		for _, file := range files {
			path := filepath.Join(uplaodDir, file.Filename)

			if err := c.SaveFile(file, path); err != nil {
				return c.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"error": err,
					})
			}

			uploadFilesLists = append(uploadFilesLists, file.Filename)

		}

		return c.JSON(fiber.Map{
			"message": "files uploded sucessfully",
			"fiels":   uploadFilesLists,
		})

	})

	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./Screenshot from 2025-12-23 17-42-22.png")
	})

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

package auth

import (
	"github.com/api-auth/models"
	"github.com/api-auth/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// type CreateUserRequest struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=3,max=32"`
// }

func AuthHandlers(route fiber.Router, db *gorm.DB) {
	route.Post("/register", func(c *fiber.Ctx) error {

		// body-parser validation
		var validate = validator.New()
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
		}

		if err := validate.Struct(&user); err != nil {
			// This returns a detailed validation error message.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// jwt hashing
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": err,
				})
		}
		user.Password = string(hashed)
		db.Create(&user)

		token, err := utils.GenerateToken(&user)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "JWT",
			Value:    token,
			HTTPOnly: !c.IsFromLocal(),
			Secure:   !c.IsFromLocal(),
			MaxAge:   3600 * 24 * 7,
		})

		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"token": token,
			})
	})

	// booking route
	route.Post("/booking", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Success message",
			"data":    "boking router",
		})
	})

}

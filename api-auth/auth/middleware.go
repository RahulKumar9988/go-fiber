package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/api-auth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookieToken := c.Cookies("JWT")
		var tokenString string

		if cookieToken != "" {
			log.Warn("token is present", cookieToken)
			tokenString = cookieToken
		} else {
			log.Warn("token is rmpty")
			authHeader := c.Get("Authorization")

			if authHeader == "" {
				log.Warn("empty token in authheader")
				return c.Status(fiber.StatusUnauthorized).
					JSON(fiber.Map{
						"status":  "failed",
						"message": "user autharized",
					})
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.Status(fiber.StatusUnauthorized).
					JSON(fiber.Map{
						"status":  "failed",
						"message": "user unautharized",
					})
			}

			tokenString = tokenParts[1]

		}

		secret := []byte("rj@secret")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
				return nil, fmt.Errorf("unxpacted signin method :%v", token.Header["alg"])
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			log.Warn("invalid token")
			c.ClearCookie("JWT")

			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{
					"eror": "invalid token",
				})
		}

		userId := token.Claims.(jwt.MapClaims)["userId"]

		if err := db.Model(&models.User{}).Where("id = ?", userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn("user not found in the DB")
			c.ClearCookie("JWT")

			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{
					"error": "unautharized",
				})
		}

		c.Locals("userId", userId)

		return c.Next()
	}
}

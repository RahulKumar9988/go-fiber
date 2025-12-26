package utils

import (
	"time"

	"github.com/api-auth/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User) (string, error) {
	secret := []byte("rj@secret")
	method := jwt.SigningMethodHS256
	claims := jwt.MapClaims{
		"userId":   user.ID,
		"username": user.Email,
		"exp":      time.Now().Add(time.Hour * 168).Unix(),
	}

	token, err := jwt.NewWithClaims(method, claims).SignedString(secret)

	if err != nil {
		return "", err
	}
	return token, nil
}

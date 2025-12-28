package models

import "github.com/api-auth/enum"

type User struct {
	ID       uint   `json:"id" gorm:"primeryKey"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Book struct {
	ID          uint            `json:"id" gorm:"primeryKey"`
	Title       string          `json:"title"`
	Status      enum.BookStatus `json:"status" gorm:"default:to_read"`
	Author      string          `json:"author"`
	Year        int             `json:"year"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	UserID      int             `json:"user_id"`
}

type UpdateBookRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
}

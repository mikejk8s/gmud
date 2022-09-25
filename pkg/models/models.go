package models

import (
	"time"

	"gorm.io/gorm"
)

type Character struct {
	Name string	`json:"name"`
	ID string	`json:"id"`
	Class string	`json:"class"`
	Race string	`json:"race"`
	Level int	`json:"level"`
	CreatedAt time.Time	`json:"created_at"`
	Alive bool	`json:"alive"`
}

type User struct {
	gorm.Model
	Name string    `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
package models

import (
	"time"
)

type Character struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Class     string    `json:"class"`
	Race      string    `json:"race"`
	Level     int       `json:"level"`
	CreatedAt time.Time `json:"created_at"`
	Alive     bool      `json:"alive"`
}

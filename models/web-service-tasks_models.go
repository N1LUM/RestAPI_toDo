package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Note struct {
	ID        uint   `gorm:"primary_key"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint   `json:"userID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

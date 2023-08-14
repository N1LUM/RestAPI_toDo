package main

import (
	"Web-service/internal/handler"
	"Web-service/internal/models"
	"github.com/jinzhu/gorm"
)

// запуск дб
func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=admin dbname=web-service-tasks password=1234 sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Миграция таблиц
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Note{})

	router := handler.SetupRouter(db)

	router.Run(":8080")
}

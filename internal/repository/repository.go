package repository

import (
	"Web-service/internal/models"
	"github.com/jinzhu/gorm"
)

type User interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}

type Note interface {
	Create()
	GetByID()
	GetAll()
	Update()
	Delete()
}

type Repository struct {
	User
	Note
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		//Note: NewNoteRepository(db),
	}
}

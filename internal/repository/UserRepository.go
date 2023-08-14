package repository

import (
	"Web-service/internal/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user *models.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) GetByID(id string) (*models.User, error) {
	var user models.User

	result := u.db.Where("id =?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserRepository) GetAll() (*[]models.User, error) {
	var users []models.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (u *UserRepository) Update(user *models.User) error {
	result := u.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) Delete(id string) error {
	result := u.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

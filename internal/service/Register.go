package service

import (
	"Web-service/internal/models"
	"Web-service/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService struct {
	rep repository.User
}

func NewUserService(rep repository.User) *UserService {
	return &UserService{rep: rep}
}

func (u *UserService) RegisterUser(c *gin.Context) {
	// созадём экземпляр и говорим что тело запроса следует сувать в user
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//хешируем
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user.Password = string(hashedPassword) // Заменяем пароль на хеш

	if err := u.rep.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
	return
}

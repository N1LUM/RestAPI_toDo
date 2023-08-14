package service

import (
	"Web-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type User interface {
	RegisterUser(c *gin.Context)
}

type Service struct {
	User
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		User: NewUserService(rep.User),
	}
}

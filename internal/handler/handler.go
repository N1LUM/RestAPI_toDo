package handler

import (
	"Web-service/internal/repository"
	"Web-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Nihao, World!",
	})
}

// handler
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)

	main := router.Group("/")
	{
		main.GET("", hello)
	}

	crud := router.Group("/crud")
	{
		crud.GET("/getAll")
		crud.GET("/getById/{id}")
		crud.POST("/create")
		crud.POST("/update")
		crud.DELETE("/delete/{id}")
	}

	join := router.Group("/join")
	{
		join.POST("/register", service.RegisterUser)
		join.POST("/login")
		join.POST("/post")
	}

	return router
}

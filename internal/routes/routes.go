package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iPhantasmic/go-gin-gorm-template/internal/services"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func InitialiseRoutes(router *gin.Engine) {
	router.GET("/health", Health)

	user := router.Group("/user")
	{
		user.GET("/", services.GetUsers)
		user.GET("/:id", services.GetUser)
		user.POST("/", services.CreateUser)
		user.PUT("/:id", services.UpdateUser)
		user.DELETE("/:id", services.DeleteUser)
	}
}

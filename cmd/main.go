package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/iPhantasmic/go-gin-gorm-template/configs"
	_ "github.com/iPhantasmic/go-gin-gorm-template/docs"
	"github.com/iPhantasmic/go-gin-gorm-template/internal/models"
	"github.com/iPhantasmic/go-gin-gorm-template/internal/routes"
)

var router *gin.Engine

func main() {
	log.Println("Server starting...")

	c, err := configs.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to initialise config: ", err)
	}

	// setup DB connection
	dbConnString := c.DBConnString
	log.Println("Database connection string: " + dbConnString)
	models.ConnectDB(dbConnString)

	// setup routes
	router = gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.InitialiseRoutes(router)

	port := c.Port
	log.Println("Running on port: " + port)
	_ = router.Run(":" + port)
}

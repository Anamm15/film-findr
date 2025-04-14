package main

import (
	"ReviewPiLem/middleware"
	"ReviewPiLem/migrations"
	"os"
	"log"

	"ReviewPiLem/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	var db *gorm.DB = config.SetUpDatabaseConnection()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := migrations.Seeder(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5454"
	}
	server.Run(":" + port)
}
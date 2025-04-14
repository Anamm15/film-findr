package main

import (
	"ReviewPiLem/middleware"
	"ReviewPiLem/migrations"
	"ReviewPiLem/repository"
	"ReviewPiLem/service"
	"ReviewPiLem/controller"
	"ReviewPiLem/routes"
	"ReviewPiLem/config"

	"os"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	var (
		db *gorm.DB = config.SetUpDatabaseConnection()
		genreRepository repository.GenreRepository = repository.NewGenreRepository(db)
		genreService service.GenreService = service.NewGenreService(genreRepository)
		genreController controller.GenreController = controller.NewGenreController(genreService)
	)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.GenreRoute(server, genreController)

	if err := migrations.Seeder(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5454"
	}
	server.Run(":" + port)
}
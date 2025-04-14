package routes

import (
	"ReviewPiLem/controller"
	"github.com/gin-gonic/gin"
)

func GenreRoute(router *gin.Engine, genreController controller.GenreController) {
	genre := router.Group("/genre")
	{
		genre.GET("/getAllGenre", genreController.GetAllGenre)
		genre.POST("/create", genreController.CreateGenre)
		genre.PUT("/update", genreController.UpdateGenre)
	}
}
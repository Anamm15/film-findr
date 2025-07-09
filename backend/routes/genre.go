package routes

import (
	"FilmFindr/controller"
	"FilmFindr/helpers"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func GenreRoute(router *gin.Engine, genreController controller.GenreController, jwtService service.JWTService) {
	genre := router.Group("/genre")
	{
		genre.GET("/getAllGenre", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.GetAllGenre)
		genre.POST("/create", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.CreateGenre)
		genre.DELETE("/delete/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.DeleteGenre)
	}
}

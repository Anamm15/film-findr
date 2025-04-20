package routes

import (
	"ReviewPiLem/controller"
	"ReviewPiLem/helpers"
	"ReviewPiLem/middleware"
	"ReviewPiLem/service"

	"github.com/gin-gonic/gin"
)

func GenreRoute(router *gin.Engine, genreController controller.GenreController, jwtService service.JWTService) {
	genre := router.Group("/genre")
	{
		genre.GET("/getAllGenre", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.GetAllGenre)
		genre.POST("/create", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.CreateGenre)
		genre.PUT("/update/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), genreController.UpdateGenre)
	}
}

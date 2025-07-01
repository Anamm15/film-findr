package routes

import (
	"FilmFindr/controller"
	"FilmFindr/helpers"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func FilmRoute(server *gin.Engine, filmController controller.FilmController, jwtService service.JWTService) {
	film := server.Group("/film")
	{
		film.GET("/getAllFilm", filmController.GetAllFilm)
		film.GET("/getFilmById/:id", filmController.GetFilmById)
		film.GET("/search", filmController.SearchFilm)
		film.POST("/create", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.CreateFilm)
		film.PUT("/update", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.UpdateFilm)
		film.PATCH("/updateStatus/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.UpdateStatus)
		film.DELETE("/delete/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.DeleteFilm)
		film.POST("/addFilmGenre", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.AddFilmGenre)
		film.DELETE("/deleteFilmGenre", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), filmController.DeleteFilmGenre)
	}
}

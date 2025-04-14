package routes

import (
	"ReviewPiLem/controller"

	"github.com/gin-gonic/gin"
)

func FilmRoute(server *gin.Engine, filmController controller.FilmController) {
	film := server.Group("/film")
	{
		film.GET("/getAllFilm", filmController.GetAllFilm)
		film.GET("/getFilmById/:id", filmController.GetFilmById)
		film.POST("/create", filmController.CreateFilm)
		film.PUT("/update", filmController.UpdateFilm)
		film.DELETE("/delete/:id", filmController.DeleteFilm)
	}
}
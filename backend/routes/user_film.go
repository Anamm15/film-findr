package routes

import (
	"ReviewPiLem/controller"
	"ReviewPiLem/middleware"
	"ReviewPiLem/service"

	"github.com/gin-gonic/gin"
)

func UserFilmRoutes(router *gin.Engine, userFilmController controller.UserFilmController, jwtService service.JWTService) {
	userFilm := router.Group("/userFilm")
	{
		userFilm.GET("/getUserFilmByUserId/:id", userFilmController.GetUserFilmByUserId)
		userFilm.POST("/create", middleware.Authenticate(jwtService), userFilmController.CreateUserFilm)
		userFilm.PATCH("/updateStatus", middleware.Authenticate(jwtService), userFilmController.UpdateStatusUserFilm)
	}
}

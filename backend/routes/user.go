package routes

import (
	"FilmFindr/controller"
	"FilmFindr/helpers"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(server *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	user := server.Group("/user")
	{
		user.GET("/getAllUser", userController.GetAllUser)
		user.GET("/:id", userController.GetUserById)
		user.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		user.POST("/register", userController.RegisterUser)
		user.POST("/login", userController.LoginUser)
		user.POST(("/logout"), middleware.Authenticate(jwtService), userController.LogoutUser)
		user.PATCH("/update", middleware.Authenticate(jwtService), userController.UpdateUser)
		user.DELETE("/delete/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), userController.DeleteUser)
	}
}

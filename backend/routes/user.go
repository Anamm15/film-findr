package routes

import (
	"ReviewPiLem/controller"
	"ReviewPiLem/helpers"
	"ReviewPiLem/middleware"
	"ReviewPiLem/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(server *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	user := server.Group("/user")
	{
		user.GET("/getAllUser", userController.GetAllUser)
		user.GET("/:id", userController.GetUserById)
		user.POST("/register", userController.RegisterUser)
		user.POST("/login", userController.LoginUser)
		user.POST(("/logout"), middleware.Authenticate(jwtService), userController.LogoutUser)
		user.PATCH("/update", middleware.Authenticate(jwtService), userController.UpdateUser)
		user.DELETE("/delete/:id", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), userController.DeleteUser)
	}
}

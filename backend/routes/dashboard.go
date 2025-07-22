package routes

import (
	"FilmFindr/controller"
	"FilmFindr/helpers"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func DashboardRoute(server *gin.Engine, dashboardController controller.DashboardController, jwtService service.JWTService) {
	dashboard := server.Group("/api/v1/dashboard")
	{
		dashboard.GET("/", middleware.Authenticate(jwtService), middleware.AuthorizeRole(helpers.ENUM_ROLE_ADMIN), dashboardController.GetDashboard)
	}
}

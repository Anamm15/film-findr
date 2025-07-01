package routes

import (
	"FilmFindr/controller"
	"FilmFindr/middleware"
	"FilmFindr/service"

	"github.com/gin-gonic/gin"
)

func ReviewRoute(server *gin.Engine, reviewController controller.ReviewController, jwtService service.JWTService) {
	review := server.Group("/review")
	{
		review.GET("/getReviewUserById/:id", reviewController.GetReviewByUserId)
		review.GET("/getReviewByFilmId/:id", reviewController.GetReviewByFilmId)
		review.POST("/create", middleware.Authenticate(jwtService), reviewController.CreateReview)
		review.PUT("/update", middleware.Authenticate(jwtService), reviewController.UpdateReview)
		review.PATCH("/updateReaksiReview", middleware.Authenticate(jwtService), reviewController.UpdateReaksiReview)
		review.DELETE("/delete/:id", middleware.Authenticate(jwtService), reviewController.DeleteReview)
	}
}

package main

import (
	"ReviewPiLem/middleware"
	"ReviewPiLem/migrations"
	"ReviewPiLem/repository"
	"ReviewPiLem/service"
	"ReviewPiLem/controller"
	"ReviewPiLem/routes"
	"ReviewPiLem/config"

	"os"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	var (
		db *gorm.DB = config.SetUpDatabaseConnection()
		jwtService service.JWTService = service.NewJWTService()

		genreRepository repository.GenreRepository = repository.NewGenreRepository(db)
		genreService service.GenreService = service.NewGenreService(genreRepository)
		genreController controller.GenreController = controller.NewGenreController(genreService)

		filmRepository repository.FilmRepository = repository.NewFilmRepository(db)
		filmService service.FilmService = service.NewFilmService(filmRepository)
		filmController controller.FilmController = controller.NewFilmController(filmService)

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService service.UserService = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)

		reviewRepository repository.ReviewRepository = repository.NewReviewRepository(db)
		reviewService service.ReviewService = service.NewReviewService(reviewRepository)
		reviewController controller.ReviewController = controller.NewReviewController(reviewService)
	)

	routes.GenreRoute(server, genreController)
	routes.FilmRoute(server, filmController)
	routes.UserRoute(server, userController, jwtService)
	routes.ReviewRoute(server, reviewController)

	if err := migrations.Seeder(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
	}
	
	port := os.Getenv("PORT")
	server.Run(":" + port)
}
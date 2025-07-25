package main

import (
	"log"
	"os"

	"FilmFindr/config"
	"FilmFindr/controller"
	"FilmFindr/middleware"
	"FilmFindr/migrations"
	"FilmFindr/repository"
	"FilmFindr/routes"
	"FilmFindr/service"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()
	server.Use(middleware.SetupCORS())

	store := cookie.NewStore([]byte("secret123"))
	store.Options(sessions.Options{
		MaxAge: 12 * 60 * 60,
	})
	server.Use(sessions.Sessions("session_token", store))

	var (
		db              *gorm.DB               = config.SetUpDatabaseConnection()
		jwtService      service.JWTService     = service.NewJWTService()
		cloudinaryCloud *cloudinary.Cloudinary = config.ConnectCloudinary()

		genreRepository        repository.GenreRepository        = repository.NewGenreRepository(db)
		filmGenreRepository    repository.FilmGenreRepository    = repository.NewFilmGenreRepository(db)
		filmGambarRepository   repository.FilmGambarRepository   = repository.NewFilmGambarRepository(db)
		userRepository         repository.UserRepository         = repository.NewUserRepository(db)
		reviewReaksiRepository repository.ReaksiReviewRepository = repository.NewReaksiReviewRepository(db)
		reviewRepository       repository.ReviewRepository       = repository.NewReviewRepository(db)
		filmRepository         repository.FilmRepository         = repository.NewFilmRepository(db)
		userFilmRepository     repository.UserFilmRepository     = repository.NewUserFilmRepository(db)

		genreService     service.GenreService     = service.NewGenreService(genreRepository)
		filmGenreService service.FilmGenreService = service.NewFilmGenreService(filmGenreRepository, db)
		userService      service.UserService      = service.NewUserService(cloudinaryCloud, userRepository)
		userFilmService  service.UserFilmService  = service.NewUserFilmService(userFilmRepository, filmRepository)
		reviewService    service.ReviewService    = service.NewReviewService(reviewRepository, reviewReaksiRepository, userFilmRepository, filmRepository)
		filmService      service.FilmService      = service.NewFilmService(db, cloudinaryCloud, filmRepository, filmGambarRepository, filmGenreRepository, reviewRepository)
		dashboardService service.DashboardService = service.NewDashboardService(filmRepository, reviewRepository, userRepository, filmService)

		genreController     controller.GenreController     = controller.NewGenreController(genreService)
		userController      controller.UserController      = controller.NewUserController(userService, jwtService)
		userFilmController  controller.UserFilmController  = controller.NewUserFilmController(userFilmService)
		reviewController    controller.ReviewController    = controller.NewReviewController(reviewService, jwtService)
		filmController      controller.FilmController      = controller.NewFilmController(filmService, filmGenreService)
		dashboardController controller.DashboardController = controller.NewDashboardController(dashboardService, jwtService)
	)

	routes.GenreRoute(server, genreController, jwtService)
	routes.FilmRoute(server, filmController, jwtService)
	routes.UserRoute(server, userController, jwtService)
	routes.ReviewRoute(server, reviewController, jwtService)
	routes.UserFilmRoutes(server, userFilmController, jwtService)
	routes.DashboardRoute(server, dashboardController, jwtService)

	if err := migrations.Seeder(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
	}

	port := os.Getenv("PORT")
	server.Run(":" + port)
}

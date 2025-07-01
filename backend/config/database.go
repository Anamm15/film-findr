package config

import (
	"fmt"
	"os"

	"FilmFindr/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := db.AutoMigrate(
		entity.User{},
		entity.Film{},
		entity.Genre{},
		entity.FilmGenre{},
		entity.FilmGambar{},
		entity.UserFilm{},
		entity.Review{},
		entity.ReaksiReview{},
	); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func ClosDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}

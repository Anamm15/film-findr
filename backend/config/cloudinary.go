package config

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func ConnectCloudinary() *cloudinary.Cloudinary {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("CLOUD_API_KEY")
	apiSecret := os.Getenv("CLOUD_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatal("Cloudinary connection error:", err)
	}

	return cld
}

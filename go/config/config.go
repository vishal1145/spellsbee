package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8000" // default port if not specified
	}
	return ":" + port
}

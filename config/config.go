package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetDBConfig() string {
	return os.Getenv("DATABASE_URL")
}

func GetRedisConfig() string {
	return os.Getenv("REDIS_URL")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetPusherConfig() (string, string, string, string, bool) {
	return os.Getenv("APP_ID"), os.Getenv("PUSHER_KEY"), os.Getenv("PUSHER_SECRET"), os.Getenv("PUSHER_CLUSTER"), os.Getenv("PUSHER_SECURE") == "true"
}

func GetPort() string {
	return os.Getenv("PORT")
}

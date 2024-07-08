package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API_KEY  string
	APP_PORT string
	DB_URL   string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		API_KEY:  getEnv("API_KEY", ""),
		APP_PORT: getEnv("APP_PORT", ":9090"),
		DB_URL:   getEnv("DB_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if _, ok := os.LookupEnv(key); ok {
		return os.Getenv(key)
	}
	return fallback
}

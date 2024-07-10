package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CHATGPT_API_KEY       string
	APP_PORT              string
	DB_URL                string
	JWT_SECRET            string
	JWT_ALGO              string
	JWT_TTL               string
	PAYSTACK_SECRET_KEY   string
	PAYSTACK_URL          string
	PAYSTACK_CALLBACK_URL string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		CHATGPT_API_KEY:       getEnv("CHATGPT_API_KEY", ""),
		APP_PORT:              getEnv("APP_PORT", ":9090"),
		DB_URL:                getEnv("DB_URL", ""),
		JWT_SECRET:            getEnv("JWT_SECRET", ""),
		JWT_ALGO:              getEnv("JWT_ALGO", ""),
		JWT_TTL:               getEnv("JWT_TTL", ""),
		PAYSTACK_SECRET_KEY:   getEnv("PAYSTACK_SECRET_KEY", ""),
		PAYSTACK_URL:          getEnv("PAYSTACK_URL", ""),
		PAYSTACK_CALLBACK_URL: getEnv("PAYSTACK_CALLBACK_URL", ""),
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

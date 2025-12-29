package config

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl   string
	DBToken string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{
		DBUrl:   os.Getenv("DB_URL"),
		DBToken: os.Getenv("DB_TOKEN"),
	}
}

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     "http://127.0.0.1:5501",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}
}

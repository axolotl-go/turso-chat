package config

import (
	"os"

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

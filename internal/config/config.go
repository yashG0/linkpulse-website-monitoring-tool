package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	Database string
	// WorkerCount   int
	// CheckInterval time.Duration
	HTTPTimeout time.Duration
}

func Load() Config {
	_ = godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Failed to load .env file")
	// }

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is required")
	}
	database := os.Getenv("DATABASE")
	if database == "" {
		log.Fatal("DATABASE environment variable is required")
	}
	return Config{
		Port:     port,
		Database: database,
	}
}

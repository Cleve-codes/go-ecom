package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	Port       string
	UploadPath string
	Env        string
}

func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "ecommerce_db"),
		JWTSecret:  getEnv("JWT_SECRET", "default-secret-change-in-production"),
		Port:       getEnv("PORT", "8082"),
		UploadPath: getEnv("UPLOAD_PATH", "./uploads"),
		Env:        getEnv("ENV", "development"),
	}

	// Validate required fields
	if config.DBPassword == "" {
		log.Fatal("DB_PASSWORD is required")
	}

	if len(config.JWTSecret) < 32 {
		log.Fatal("JWT_SECRET must be at least 32 characters long")
	}

	return config
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

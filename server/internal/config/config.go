package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port        string
    Environment string
    DatabaseURL string
    JWTSecret   string
    LogLevel    string
    RedisURL    string
}

func Load() *Config {
    // Load .env file in development
    if os.Getenv("ENVIRONMENT") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("No .env file found")
        }
    }

    return &Config{
        Port:        getEnv("PORT", "8080"),
        Environment: getEnv("ENVIRONMENT", "development"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost/bitbox_db?sslmode=disable"),
        JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
        LogLevel:    getEnv("LOG_LEVEL", "info"),
        RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}


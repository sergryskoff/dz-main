// package config рабочие параметры
package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr string
	Db         DbConfig
}

type DbConfig struct {
	Dsn string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file, using default config")
	}
	if os.Getenv("PORT_SERVER") == "" || os.Getenv("DB_HOST") == "" || os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_NAME") == "" || os.Getenv("DB_PORT") == "" {
		slog.Info("Can't load config from .env file, using default config")
	}
	return &Config{
		ServerAddr: getEnv("PORT_SERVER", ":8081"),
		Db: DbConfig{
			Dsn: fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
				getEnv("DB_HOST", "localhost"),
				getEnv("DB_USER", "postgres"),
				getEnv("DB_PASSWORD", "my_pass"),
				getEnv("DB_NAME", "product"),
				getEnv("DB_PORT", "5432"),
			),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

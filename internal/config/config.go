package config

import (
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/scdb/server/internal/logger"
)

type Config struct {
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBMaxConnections string
	DBConnTimeout    string
	DBSSLMode        string
	WEBUrl           string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		logger.Warning(".env файл не найден, пробую использовать системные переменные")
	}

	AppConfig = &Config{
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPassword:       getEnv("DB_PASSWORD", "password"),
		DBName:           getEnv("DB_NAME", "db"),
		DBMaxConnections: getEnv("DB_MAX_CONNECTIONS", "100"),
		DBConnTimeout:    getEnv("DB_CONN_TIMEOUT", "30"),
		DBSSLMode:        getEnv("DB_SSL_MODE", "require"),
		WEBUrl:           getEnv("WEB_URL", ""),
	}

	logger.Success("Конфигурация env загружена")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

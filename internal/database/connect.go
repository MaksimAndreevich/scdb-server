package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/scdb/server/internal/config"
	"gitlab.com/scdb/server/internal/logger"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {

	// Формируем строку для подключения к pq
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
	)

	db, err := sql.Open("postgres", dsn)

	DB = db

	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных:", err)
	}

	if err := DB.Ping(); err != nil {
		logger.Fatal(`База данных не отвечает: `, err)

	}

	logger.Success("Успешное подключение к базе данных!")

	return DB, nil
}

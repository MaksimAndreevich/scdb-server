package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/scdb/server/internal/config"
	"github.com/scdb/server/internal/logger"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {

	// Формируем строку для подключения к pq
	dsn := fmt.Sprintf(
		// TODO: Добавить sslmode=require когда будет готов сертификат
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Настройка пула соединений
	maxConns, err := strconv.Atoi(config.AppConfig.DBMaxConnections)
	if err != nil {
		logger.Fatal("Ошибка парсинга DB_MAX_CONNECTIONS:", err)
	}
	db.SetMaxOpenConns(maxConns)

	connTimeout, err := strconv.Atoi(config.AppConfig.DBConnTimeout)
	if err != nil {
		logger.Fatal("Ошибка парсинга DB_CONN_TIMEOUT:", err)
	}
	db.SetConnMaxLifetime(time.Duration(connTimeout) * time.Second)

	DB = db

	if err := DB.Ping(); err != nil {
		logger.Fatal(`База данных не отвечает: `, err)

	}

	logger.Success("Успешное подключение к базе данных!")

	return DB, nil
}

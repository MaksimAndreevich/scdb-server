# SCDB Server

Серверная часть системы управления образовательными организациями (SCDB - School Control Database).

## Описание

SCDB Server - это REST API сервер, предоставляющий доступ к базе данных образовательных организаций. Сервер реализован на Go с использованием Gin framework и PostgreSQL в качестве базы данных.

## Основные возможности

- Получение информации об образовательных организациях
- Поиск организаций с фильтрацией и пагинацией
- Получение статистики по базам данных
- Поддержка различных типов образовательных учреждений

## Технический стек

- Go 1.24.3
- Gin Web Framework
- PostgreSQL 17.4
- Docker & Docker Compose

## Требования

- Go 1.24.3 или выше
- Docker и Docker Compose
- PostgreSQL 17.4 (если запуск без Docker)

## Установка и запуск

### Использование Docker

1. Клонируйте репозиторий:

```bash
git clone https://gitlab.com/scdb/server.git
cd server
```

2. Запустите приложение с помощью Docker Compose:

```bash
docker-compose up -d
```

Сервер будет доступен по адресу: http://localhost:8080

### Локальная установка

1. Установите зависимости:

```bash
go mod download
```

2. Создайте файл .env в корневой директории:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=db
```

3. Запустите сервер:

```bash
go run main.go
```

## API Endpoints

### Организации

- `GET /api/organizations/{id}` - Получение организации по ID
- `GET /api/organizations` - Получение списка организаций с пагинацией и фильтрацией
  - Параметры:
    - search - Поисковый запрос
    - federal_district_id - ID федерального округа
    - region_id - ID региона
    - city_id - ID города
    - education_type_key - Тип образования
    - page - Номер страницы
    - per_page - Количество записей на странице

### Статистика

- `GET /api/stats` - Получение общей статистики по базам данных

## Разработка

### Структура проекта

```
.
├── docs/                    # Документация
├── internal/               # Внутренний код приложения
│   ├── config/            # Конфигурация
│   ├── database/          # Работа с базой данных
│   └── routers/           # Маршрутизация
├── main.go                # Точка входа
├── Dockerfile             # Конфигурация Docker
├── docker-compose.yml     # Конфигурация Docker Compose
└── go.mod                 # Зависимости Go
```

### Hot Reload

Для разработки используется Air для автоматической перезагрузки при изменении кода:

```bash
air
```

# 1. Этап сборки
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum, чтобы использовать кеширование зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем бинарник для Linux (если нужно, можно указать конкретное имя)
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Указываем, какой порт слушать
EXPOSE 8080

# Команда запуска
CMD ["./server"]

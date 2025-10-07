# Стадия 1: Сборка (Build stage)
FROM golang:1.18-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей для кеширования зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем приложение в бинарный файл
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o golang-auth .

# Стадия 2: Runtime
FROM alpine:latest

# Устанавливаем ca-certificates для HTTPS запросов
RUN apk --no-cache add ca-certificates

# Создаем непривилегированного пользователя
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем бинарный файл из стадии сборки
COPY --from=builder /app/golang-auth .

# Меняем владельца файла
RUN chown appuser:appgroup golang-auth

# Переключаемся на непривилегированного пользователя
USER appuser

# Открываем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD ["./golang-auth"]

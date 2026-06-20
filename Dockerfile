# Оптимизированный многоэтапный Dockerfile для Go-приложения
FROM golang:1.22-alpine AS builder

# Настройка окружения
WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Копируем только зависимости для лучшего кэширования
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарный файл
RUN go build -ldflags="-s -w" -o main .

# Минимальная базовая imagem для продакшена
FROM alpine:3.19

# Установка необходимых утилит для работы с контейнерами и времени
RUN apk --no-cache add ca-certificates tzdata

# Создание пользователя и группы (безопасность)
RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

WORKDIR /app

# Копируем скомпилированный бинарный файл из стадии сборки
COPY --from=builder --chown=appuser:appgroup /app/main .
COPY --from=builder --chown=appuser:appgroup /app/*.json ./config/

# Устанавливает владельца файлов контейнера
RUN chown -R appuser:appgroup /app

USER appuser

# Добавляет локальное время для корректной работы таймеров в приложении
ENV TZ=Europe/Moscow

EXPOSE 8080

# Запуск приложения из образа сборки через shell с проверкой окружения
ENTRYPOINT ["/app/main"]
# Используем Go 1.21 как базовый образ (или другую версию по необходимости)
FROM golang:1.21-alpine AS builder

# Установка необходимых пакетов
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Копируем go.mod и go.sum (если есть) для быстрой сборки модулей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Переключаемся на версию Go 1.21-alpine для финального образа с меньшим размером
FROM golang:1.21-alpine

# Установка необходимых системных пакетов
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Копируем go.mod и go.sum из builder-стадии
COPY --from=builder /app/go.mod .
COPY --from=builder /app/go.sum .

# Запускаем сборку бота на Go
COPY --from=builder /app/bot ./bot
RUN CGO_ENABLED=0 GOOS=linux go build -o bot ./bot

# Установка timezone (опционально, но рекомендуется)
ENV TZ=UTC
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

# Копируем config.json (если существует в репозитории или нужно смонтировать извне)
COPY config.json . 2>/dev/null || true

# Создаём необходимые директории и устанавливаем права
RUN mkdir -p /data /tmp && \
    chown -R nobody:nogroup /app /data /tmp

# Запускаем бот (без привилегий пользователя root)
USER nobody

EXPOSE 8080

ENTRYPOINT ["./bot"]
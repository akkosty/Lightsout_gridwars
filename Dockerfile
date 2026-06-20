# Базовый образ с Go
FROM golang:1.21-alpine AS builder

# Установка зависимостей для компиляции в alpine
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Копируем go.mod и go.sum (если будут созданы)
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем с целевой платформой для Docker
ARG TARGETPLATFORM
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(if [ "$TARGETPLATFORM" = "amd64" ] || [ "$TARGETPLATFORM" = "x86_64" ]; then echo "amd64"; else echo "arm64"; fi) go build -ldflags="-s -w" -o bot/main .

# Финальный этап: минимальный образ
FROM alpine:latest

WORKDIR /app

# Копируем скомпилированный бинарный файл
COPY --from=builder /app/bot/main .

# Добавляем зависимости для работы (если необходимы)
RUN apk add --no-cache curl

# Переменная окружения для токена
ENV TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN}

EXPOSE 8080

CMD ["/app/bot/main"]
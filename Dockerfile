# Используем официальный Go образ для Linux
FROM golang:1.21-alpine AS builder

# Установка необходимых пакетов
RUN apk add --no-cache ca-certificates tzdata git

WORKDIR /app

# Копируем go.mod для быстрой сборки модулей
COPY go.mod ./
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Собираем бинарный файл
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Минимальный образ для запуска
FROM alpine:latest

# Установка системных зависимостей и времени
RUN apk add --no-cache ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

# Установка timezone
ENV TZ=UTC
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

# Копируем скомпилированный бинарный файл из стадии builder
COPY --from=builder /app/main .

# Создаём необходимые директории для телеграм бота (если используются)
RUN mkdir -p /tmp

# Записываем конфиг из образа или создаём пустой файл
RUN touch config.json || true

# Права доступа
RUN chown -R nobody:nogroup /app /tmp && \
    chmod +x main

USER nobody

EXPOSE 8080

CMD ["./main"]
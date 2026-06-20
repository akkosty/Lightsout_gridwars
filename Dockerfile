# Используем официальный Go образ для Linux
FROM golang:1.21-alpine AS builder

# Установка необходимых пакетов
RUN apk add --no-cache ca-certificates tzdata git

WORKDIR /app

# Копируем только необходимые файлы проекта (без go.sum)
COPY main.go ./main.go
COPY go.mod ./go.mod  # Важно: нужен для управления зависимостями

# Копируем конфиг, если существует
COPY config.json ./config.json || true

# Запускаем обновление модулей и скачивание зависимостей
RUN go mod download

# Собираем бинарный файл без CGO
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

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

# Создаём необходимые директории и временную зону
RUN mkdir -p /tmp

# Права доступа
RUN chown -R nobody:nogroup /app /tmp && \
    chmod +x main

USER nobody

EXPOSE 8080

CMD ["./main"]
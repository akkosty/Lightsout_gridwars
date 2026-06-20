# Использовать официальный образ Go для разработки
FROM golang:1.21-alpine AS builder

# Установить зависимости и скомпилировать бот
WORKDIR /app
COPY . .
RUN go mod download && go build -o main .

# Минимальный прогмаобраз для продакшена
FROM alpine:3.19
LABEL maintainer="akkosty"
WORKDIR /root/

# Копируем только скомпилированный бинарный файл из стадии сборки
COPY --from=builder /app/main .
COPY --from=builder /app/*.json ./config/

# Добавляем скрипт запуска
COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/docker-entrypoint.sh"]
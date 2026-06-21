# Базовый образ Go
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod .
RUN apk --no-cache add git && go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Минимальный образ с Go ботом
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/go.mod .

ENV TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN}
CMD ["/main"]
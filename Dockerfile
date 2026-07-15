# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY bot/go.mod bot/go.sum ./

RUN go mod download

COPY bot/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o bot .

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/bot ./bot

COPY img/ ./img/

ENV TELEGRAM_BOT_TOKEN=""

CMD ["./bot"]

# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /src

# Cache go.mod dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bot ./cmd

# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Copy the compiled binary from builder
COPY --from=builder /app/bot .

# Environment variable for bot token (set by Render)
ENV TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN}

CMD ["./bot"]

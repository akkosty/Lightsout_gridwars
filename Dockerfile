# Using official Go image for Linux
FROM golang:1.21-alpine AS builder

# Install necessary packages
RUN apk add --no-cache ca-certificates tzdata git

WORKDIR /app

# Copy only necessary project files (go.sum is critical for dependencies)
COPY main.go ./main.go
COPY go.mod ./go.mod
COPY go.sum ./go.sum

# Copy config if exists
COPY config.json ./config.json || true

# Run module download and dependency fetching
RUN go mod download

# Build binary without CGO
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Minimal image for running
FROM alpine:latest

# Install system dependencies and timezone
RUN apk add --no-cache ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

# Setup timezone
ENV TZ=UTC
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

# Copy compiled binary from builder stage
COPY --from=builder /app/main .

# Create necessary directories and temp zone
RUN mkdir -p /tmp

# Permissions
RUN chown -R nobody:nogroup /app /tmp && \
    chmod +x main

USER nobody

EXPOSE 8080

CMD ["/app/main"]
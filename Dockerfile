# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /src

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build from bot directory
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bot ./bot

# Runtime stage
FROM alpine:3.19
WORKDIR /app

# Copy binary and assets
COPY --from=builder /app/bot .
COPY img ./img

CMD ["./bot"]
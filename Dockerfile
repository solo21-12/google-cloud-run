# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/gcr-api ./cmd/main.go

# Stage 2: Set up the runtime environment
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/gcr-api .

# Install supervisor
RUN apk add --no-cache supervisor

# Copy the supervisord configuration file
COPY supervisord.conf /etc/supervisord.conf

# Expose the application port
EXPOSE 8081

# Use supervisord to manage the application process
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]

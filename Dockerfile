FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/gcr-api ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/gcr-api .
EXPOSE 8081
CMD ["./gcr-api"]

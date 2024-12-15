# Stage 1: Builder
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/main.go

# Stage 2: Final runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
RUN chmod +x /app/app

EXPOSE 1234 2222

CMD ["./app"]

# Use the official Golang image as the build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o mutual-fund-engine ./cmd/server

# Use a minimal image for the runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mutual-fund-engine .
COPY internal/config/config.yaml ./internal/config/config.yaml
EXPOSE 8080
CMD ["./mutual-fund-engine"]

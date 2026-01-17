# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git for fetching dependencies
RUN apk add --no-cache git

# Download dependencies first (cached if go.mod/go.sum don't change)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
# Build the binary, disabling CGO for a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -v -o api-server ./cmd/main.go

# Run Stage
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/api-server .
# Copy any other necessary files (e.g., config, schema.sql if needed at runtime)
# COPY --from=builder /app/config ./config 

# Expose port (documentary)
EXPOSE 8080

# Run the binary
CMD ["./api-server"]

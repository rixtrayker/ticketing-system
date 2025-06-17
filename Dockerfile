# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN make build

# Final stage
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/server .

# Copy migrations
COPY --from=builder /app/migrations ./migrations

# Copy config files
COPY --from=builder /app/.env.example ./.env

# Expose port
EXPOSE 8080

# Run the application
CMD ["./server"] 
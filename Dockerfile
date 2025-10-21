# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies (gcc, musl-dev for CGO)
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build all binaries with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/bin/reader ./cmd/reader/main.go
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/bin/writer ./cmd/writer/main.go

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache sqlite-libs ca-certificates

WORKDIR /app

# Copy binaries from builder
COPY --from=builder /app/bin/reader /app/reader
COPY --from=builder /app/bin/writer /app/writer

# Copy .env file
COPY .env* ./

# Create uploads directory
RUN mkdir -p /app/uploads

# Expose port
EXPOSE 8080

# Default command (can be overridden by SCOPE)
CMD ["sh", "-c", "if [ \"$SCOPE\" = \"writer\" ]; then ./writer; else ./reader; fi"]



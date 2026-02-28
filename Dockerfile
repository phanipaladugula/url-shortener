# STAGE 1: Build the binary
FROM golang:1.24-alpine AS builder

# Install build tools
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy dependency files first (for better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build a statically linked binary (perfect for Alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener .

# STAGE 2: Run the binary
FROM alpine:latest  

# Add security: Don't run as root in production
RUN adduser -D gouser
USER gouser

WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/url-shortener .

# Expose the port your server listens on
EXPOSE 8080

# Run the application
CMD ["./url-shortener"]
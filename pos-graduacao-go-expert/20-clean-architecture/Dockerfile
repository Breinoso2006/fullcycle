# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Install git for go modules
RUN apk add --no-cache git

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN cd cmd/ordersystem && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/cmd/ordersystem/main .
COPY --from=builder /app/cmd/ordersystem/.env .

# Expose the ports your app uses
EXPOSE 8000 50051 8080

# Command to run
CMD ["./main"]

# Stage 1: Build application
FROM golang:1.24.0 AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go build -o todolist ./cmd/main.go

# Stage 2: Run application
FROM debian:latest

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/todolist .

# Ensure the binary has execution permissions
RUN chmod +x ./todolist

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./todolist"]

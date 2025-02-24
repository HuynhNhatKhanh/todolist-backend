# Use the official Golang image as the build environment
FROM golang:1.24.0 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project into the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application binary
RUN go build -o todolist ./cmd/main.go

# Use a minimal Debian-based image to run the application
FROM debian:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/todolist .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./todolist"]

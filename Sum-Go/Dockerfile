# Step 1: Use a Golang base image to build the Go app
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first (if they exist)
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the rest of the Go application source code
COPY . .

# Build the Go application
RUN go build -o myapp .

# Step 2: Create a minimal runtime image (distroless base image or alpine)
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/myapp .

# Optionally copy any static assets (public folder in this case)
COPY ./public /root/public

# Expose port 8080 (as your application is running on this port)
EXPOSE 8080

# Command to run the Go application when the container starts
CMD ["./myapp"]

# Step 1: Build the Go application
FROM golang:1.18-alpine as builder

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o myapp .

# Step 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/myapp .

# Command to run when the container starts
CMD ["./myapp"]

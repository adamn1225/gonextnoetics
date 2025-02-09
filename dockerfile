# Use a small Golang image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy everything into the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Expose the port for the Go server
EXPOSE 5000

# Run the Go app
CMD ["./main"]

# Use a minimal Go image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source files into the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go binary
RUN go build -o backend

# Expose port 5000 (since we changed it from 8080)
EXPOSE 5000

# Run the compiled Go application
CMD ["./backend"]

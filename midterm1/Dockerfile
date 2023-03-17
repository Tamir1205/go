FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go app
RUN go build -o main cmd/api/main.go

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["sh", "-c", "sleep 10 && /app/main"]
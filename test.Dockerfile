# Stage 1: Build the Go application
FROM golang:1.22.8-alpine

WORKDIR /app

# Copy the application code into the container
COPY . .

# Set the working directory to where `go.mod` is located
WORKDIR /app/src

# Download dependencies and build the Go application
RUN go mod tidy

CMD ["go", "test", "-v", "./..."] 

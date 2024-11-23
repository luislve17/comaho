# Stage 1: Build the Go application
FROM golang:1.22.8-alpine

WORKDIR /app

# Copy the application code into the container
COPY . .

# Set the working directory to where `go.mod` is located
WORKDIR /app/src

# Download dependencies and build the Go application
RUN go mod tidy

# Create a non-root user and switch to it
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Fix permissions for the Go module cache and related directories
RUN mkdir -p /go/pkg/mod && chown -R appuser:appgroup /go

# Switch to non-root user
USER appuser

CMD ["go", "test", "-failfast", "./..."]


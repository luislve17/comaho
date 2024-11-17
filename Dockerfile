# Stage 1: Build the Go application
FROM golang:1.22.8-alpine AS builder

WORKDIR /app

# Copy the application code into the container (including frontend)
COPY . .

# Set the working directory to where `go.mod` is located
WORKDIR /app/src

# Download dependencies and build the Go application
RUN go mod tidy
RUN go build -o comaho .

# Stage 2: Final image to run the application
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/src/comaho .

# Copy the frontend templates (index.html)
COPY --from=builder /app/src/templates /app/templates

# Expose the default port (this can be overridden by an environment variable)
EXPOSE 8080

# Set the environment variable COMAHO_PORT to the default value
ENV COMAHO_PORT=8080

# Command to run the Go application
CMD ["./comaho"]


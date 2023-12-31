# Use Go 1.21 as the base image
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /workspace

# Copy go mod and sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the application targeting the main.go in the cmd directory
RUN CGO_ENABLED=0 GOOS=linux go build -o app_gateway ./main.go

# Start a new stage with a minimal image for smaller size
FROM alpine:latest

# Add ca-certificates for secure connections and bash for the entrypoint script
RUN apk --no-cache add ca-certificates bash netcat-openbsd curl

# Copy the binary from the builder stage to the current stage
COPY --from=builder /workspace/app_gateway /app_gateway
COPY --from=builder /workspace/configs/config.yaml /configs/config.yaml

# Expose any necessary ports (if required)
EXPOSE 8080

# Command to run the application
CMD ["./app_gateway"]

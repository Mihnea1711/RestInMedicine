#!/bin/bash

echo "[RABBIT_ENTRYPOINT] Starting entry script..."

# Function to wait for RabbitMQ to be ready
wait_for_rabbit() {
    echo "[RABBIT_ENTRYPOINT] Waiting for RabbitMQ to be ready..."
    retries=10
    while ((retries > 0)); do
        nc -z rabbitmq 5672 && echo "[RABBIT_ENTRYPOINT] RabbitMQ is ready." && return 0
        ((retries--))
        sleep 5
    done
    echo "[RABBIT_ENTRYPOINT] Failed to connect to RabbitMQ after retries."
    exit 1
}

# Execute the wait function
wait_for_rabbit

# Start your Go application
echo "[RABBIT_ENTRYPOINT] Starting the application..."
/app_rabbit


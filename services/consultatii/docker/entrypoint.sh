#!/bin/bash

echo "[CONSULTATION_ENTRYPOINT] Starting entry script..."

# Function to wait for MongoDB to be ready
wait_for_mongodb() {
    echo "[CONSULTATION_ENTRYPOINT] Waiting for MongoDB to be ready..."
    while true; do
        nc -z consultation_mongodb 27017 && echo "[CONSULTATION_ENTRYPOINT] MongoDB is ready." && break
    done
}

wait_for_redis() {
    echo "[CONSULTATION_ENTRYPOINT] Waiting for redis to be ready..."
    while true; do
        nc -z consultation_redis 6379 && echo "[CONSULTATION_ENTRYPOINT] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mongodb
wait_for_redis

# Start your Go application
echo "[CONSULTATION_ENTRYPOINT] Starting the application..."
/app_consultatii

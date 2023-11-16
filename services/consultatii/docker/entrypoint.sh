#!/bin/bash

echo "[CONSULTATION] Starting entry script..."

# Function to wait for MongoDB to be ready
wait_for_mongodb() {
    echo "[CONSULTATION] Waiting for MongoDB to be ready..."
    while true; do
        nc -z mongodb 27017 && echo "[CONSULTATION] MongoDB is ready." && break
    done
}

wait_for_redis() {
    echo "[CONSULTATION] Waiting for redis to be ready..."
    while true; do
        nc -z redis 6379 && echo "[CONSULTATION] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mongodb
wait_for_redis

# Start your Go application
echo "[CONSULTATION] Starting the application..."
/app_consultatii

#!/bin/bash

echo "[CONSULTATII] Starting entry script..."

# Function to wait for MongoDB to be ready
wait_for_mongodb() {
    echo "[CONSULTATII] Waiting for MongoDB to be ready..."
    while true; do
        nc -z mongodb 27017 && echo "[CONSULTATII] MongoDB is ready." && break
    done
}

wait_for_redis() {
    echo "[CONSULTATII] Waiting for redis to be ready..."
    while true; do
        nc -z redis 6379 && echo "[CONSULTATII] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mongodb
wait_for_redis

# Start your Go application
echo "[CONSULTATII] Starting the application..."
/app_consultatii

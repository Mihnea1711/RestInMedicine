#!/bin/bash

echo "[PROGRAMARE] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[PROGRAMARE] Waiting for mysql to be ready..."
    while true; do
        nc -z mysql 3306 && echo "[PROGRAMARE] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[PROGRAMARE] Waiting for redis to be ready..."
    while true; do
        nc -z redis 6379 && echo "[PROGRAMARE] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
wait_for_redis

# Start your Go application
echo "[PROGRAMARE] Starting the application..."
/app_programari

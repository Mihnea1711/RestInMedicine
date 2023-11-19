#!/bin/bash

echo "[DOCTOR] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[DOCTOR] Waiting for mysql to be ready..."
    while true; do
        nc -z doctor_mysql 3306 && echo "[DOCTOR] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[DOCTOR] Waiting for redis to be ready..."
    while true; do
        nc -z doctor_redis 6379 && echo "[DOCTOR] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
wait_for_redis

# Start your Go application
echo "[DOCTOR] Starting the application..."
/app_doctori

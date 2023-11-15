#!/bin/bash

echo "[PATIENT] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[PATIENT] Waiting for mysql to be ready..."
    while true; do
        nc -z mysql 3306 && echo "[PATIENT] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[PATIENT] Waiting for redis to be ready..."
    while true; do
        nc -z redis 6379 && echo "[PATIENT] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
wait_for_redis

# Start your Go application
echo "[PATIENT] Starting the application..."
/app_pacienti

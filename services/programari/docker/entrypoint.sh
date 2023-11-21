#!/bin/bash

echo "[APPOINTMENT] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[APPOINTMENT] Waiting for mysql to be ready..."
    while true; do
        nc -z appointment_mysql 3306 && echo "[APPOINTMENT] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[APPOINTMENT] Waiting for redis to be ready..."
    while true; do
        nc -z appointment_redis 6379 && echo "[APPOINTMENT] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
wait_for_redis

# Start your Go application
echo "[APPOINTMENT] Starting the application..."
/app_programari

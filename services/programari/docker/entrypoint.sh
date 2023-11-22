#!/bin/bash

echo "[APPOINTMENT_ENTRYPOINT] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[APPOINTMENT_ENTRYPOINT] Waiting for mysql to be ready..."
    while true; do
        nc -z appointment_mysql 3306 && echo "[APPOINTMENT_ENTRYPOINT] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[APPOINTMENT_ENTRYPOINT] Waiting for redis to be ready..."
    while true; do
        nc -z appointment_redis 6379 && echo "[APPOINTMENT_ENTRYPOINT] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
wait_for_redis

# Start your Go application
echo "[APPOINTMENT_ENTRYPOINT] Starting the application..."
/app_programari

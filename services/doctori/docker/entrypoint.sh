#!/bin/bash

echo "[DOCTOR_ENTRYPOINT] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[DOCTOR_ENTRYPOINT] Waiting for mysql to be ready..."
    while true; do
        nc -z pdp_mysql 3306 && echo "[DOCTOR_ENTRYPOINT] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[DOCTOR_ENTRYPOINT] Waiting for redis to be ready..."
    while true; do
        nc -z doctor_redis 6379 && echo "[DOCTOR_ENTRYPOINT] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
wait_for_redis

# Start your Go application
echo "[DOCTOR_ENTRYPOINT] Starting the application..."
/app_doctori

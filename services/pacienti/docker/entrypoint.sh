#!/bin/bash

echo "[PATIENT_ENTRYPOINT] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[PATIENT_ENTRYPOINT] Waiting for MySQL to be ready..."
    while true; do
        nc -z patient_mysql 3306 && echo "[PATIENT_ENTRYPOINT] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[PATIENT_ENTRYPOINT] Waiting for Redis to be ready..."
    while true; do
        nc -z patient_redis 6379 && echo "[PATIENT_ENTRYPOINT] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
echo "[PATIENT_ENTRYPOINT] After waiting for MySQL. Continuing..."
wait_for_redis
echo "[PATIENT_ENTRYPOINT] After waiting for Redis. Continuing..."

# Start your Go application
echo "[PATIENT_ENTRYPOINT] Starting the application..."
/app_pacienti

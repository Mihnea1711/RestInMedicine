#!/bin/bash

echo "[IDM] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[IDM] Waiting for mysql to be ready..."
    while true; do
        nc -z mysql 3306 && echo "[IDM] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[IDM] Waiting for redis to be ready..."
    while true; do
        nc -z redis 6379 && echo "[IDM] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
wait_for_redis

# Start your Go application
echo "[IDM] Starting the application..."
/app_idm

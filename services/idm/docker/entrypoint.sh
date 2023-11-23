#!/bin/bash

echo "[IDM_ENTRYPOINT] Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "[IDM_ENTRYPOINT] Waiting for mysql to be ready..."
    while true; do
        nc -z idm_mysql 3306 && echo "[IDM_ENTRYPOINT] MySQL is ready." && break
    done
}

wait_for_redis() {
    echo "[IDM_ENTRYPOINT] Waiting for redis to be ready..."
    while true; do
        nc -z idm_redis 6379 && echo "[IDM_ENTRYPOINT] Redis is ready." && break
    done
}

# Execute the wait functions
wait_for_mysql
wait_for_redis

# Start your Go application
echo "[IDM_ENTRYPOINT] Starting the application..."
/app_idm

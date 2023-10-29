#!/bin/bash

echo "Starting entry script..."

# Function to wait for MySQL to be ready
wait_for_mysql() {
    echo "Waiting for mysql to be ready..."
    while true; do
        nc -z mysql 3306 && echo "MySQL is ready." && break
    done
}

# Execute the wait function
wait_for_mysql

# Start your Go application
echo "Starting the application..."
/app_doctori

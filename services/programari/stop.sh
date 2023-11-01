#!/bin/bash

echo "[PROGRAMARE] Stopping existing containers..."
docker compose down

echo "[PROGRAMARE] Removing MySQL data volume..."
docker volume rm programari_mysql_data

echo "[PROGRAMARE] Removing Redis data volume..."
docker volume rm programari_redis_data

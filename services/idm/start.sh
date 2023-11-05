#!/bin/bash

echo "[PROGRAMARE] Stopping existing containers..."
docker compose down

echo "[PROGRAMARE] Removing MySQL data volume..."
docker volume rm idm_mysql_data

echo "[PROGRAMARE] Removing Redis data volume..."
docker volume rm idm_redis_data

echo "[PROGRAMARE] Building Docker images..."
docker compose build

echo "[PROGRAMARE] Starting containers..."
docker compose up --force-recreate

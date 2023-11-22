#!/bin/bash

echo "[APPOINTMENT] Stopping existing containers..."
docker compose down

echo "[APPOINTMENT] Removing unused volumes..."
docker volume prune --force

echo "[APPOINTMENT] Removing MySQL data volume..."
docker volume rm programari_mysql_data

echo "[APPOINTMENT] Removing Redis data volume..."
docker volume rm programari_redis_data

echo "[APPOINTMENT] Removing unused networks..."
docker network prune --force

echo "[APPOINTMENT] Removing unused images..."
docker image prune --force

echo "[APPOINTMENT] Building Docker images..."
docker compose build

echo "[APPOINTMENT] Starting containers..."
docker compose up --force-recreate

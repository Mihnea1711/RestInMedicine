#!/bin/bash

echo "[DOCTOR] Stopping existing containers..."
docker compose down

echo "[PATIENT] Removing unused volumes..."
docker volume prune --force

echo "[DOCTOR] Removing MySQL data volume..."
docker volume rm doctori_mysql_data

echo "[DOCTOR] Removing Redis data volume..."
docker volume rm doctori_redis_data

echo "[PATIENT] Removing unused networks..."
docker network prune --force

echo "[PATIENT] Removing unused images..."
docker image prune --force

echo "[DOCTOR] Building Docker images..."
docker compose build

echo "[DOCTOR] Starting containers..."
docker compose up --force-recreate

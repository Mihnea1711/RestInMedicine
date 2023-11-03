#!/bin/bash

echo "[CONSULTATIE] Stopping existing containers..."
docker compose down

echo "[CONSULTATIE] Removing MySQL data volume..."
docker volume rm consultatii_mysql_data

echo "[CONSULTATIE] Removing Redis data volume..."
docker volume rm consultatii_redis_data

echo "[CONSULTATIE] Building Docker images..."
docker compose build

echo "[CONSULTATIE] Starting containers..."
docker compose up --force-recreate

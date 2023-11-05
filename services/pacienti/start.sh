#!/bin/bash

echo "[PACIENT] Stopping existing containers..."
docker compose down

echo "[PACIENT] Removing MySQL data volume..."
docker volume rm pacienti_mysql_data

echo "[PACIENT] Removing Redis data volume..."
docker volume rm pacienti_redis_data

echo "[PACIENT] Building Docker images..."
docker compose build

echo "[PACIENT] Starting containers..."
docker compose up --force-recreate

#!/bin/bash

echo "[PATIENT] Stopping and removing existing containers..."
docker compose down

echo "[PATIENT] Removing unused volumes..."
docker volume prune --force

echo "[PATIENT] Removing MySQL data volume..."
docker volume rm pacienti_mysql_data

echo "[PATIENT] Removing Redis data volume..."
docker volume rm pacienti_redis_data

echo "[PATIENT] Removing unused networks..."
docker network prune --force

echo "[PATIENT] Removing unused images..."
docker image prune --force

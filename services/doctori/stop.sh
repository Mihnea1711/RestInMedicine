#!/bin/bash

echo "[DOCTOR] Stopping existing containers..."
docker compose down

echo "[DOCTOR] Removing unused volumes..."
docker volume prune --force

echo "[DOCTOR] Removing MySQL data volume..."
docker volume rm doctori_mysql_data

echo "[DOCTOR] Removing Redis data volume..."
docker volume rm doctori_redis_data

echo "[DOCTOR] Removing unused networks..."
docker network prune --force

echo "[DOCTOR] Removing unused images..."
docker image prune --force

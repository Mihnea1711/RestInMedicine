#!/bin/bash

echo "[IDM] Stopping existing containers..."
docker compose down

echo "[APPOINTMENT] Removing unused volumes..."
docker volume prune --force

echo "[IDM] Removing MySQL data volume..."
docker volume rm idm_mysql_data

echo "[IDM] Removing Redis data volume..."
docker volume rm idm_redis_data

echo "[APPOINTMENT] Removing unused networks..."
docker network prune --force

echo "[APPOINTMENT] Removing unused images..."
docker image prune --force

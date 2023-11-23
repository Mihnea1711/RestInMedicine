#!/bin/bash

echo "[IDM] Stopping existing containers..."
docker compose down

echo "[IDM] Removing unused volumes..."
docker volume prune --force

echo "[IDM] Removing MySQL data volume..."
docker volume rm idm_mysql_data

echo "[IDM] Removing Redis data volume..."
docker volume rm idm_redis_data

echo "[IDM] Removing unused networks..."
docker network prune --force

echo "[IDM] Removing unused images..."
docker image prune --force

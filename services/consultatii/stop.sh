#!/bin/bash

echo "[CONSULTATION] Stopping existing containers..."
docker compose down

echo "[CONSULTATION] Removing unused volumes..."
docker volume prune --force

echo "[CONSULTATION] Removing MySQL data volume..."
docker volume rm consultatii_mysql_data

echo "[CONSULTATION] Removing Redis data volume..."
docker volume rm consultatii_redis_data

echo "[CONSULTATION] Removing unused networks..."
docker network prune --force

echo "[CONSULTATION] Removing unused images..."
docker image prune --force
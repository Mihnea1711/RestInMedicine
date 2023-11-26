#!/bin/bash

echo "[APPOINTMENT] Stopping existing containers..."
docker compose down

echo "[APPOINTMENT] Removing unused volumes..."
docker volume prune --force

echo "[APPOINTMENT] Removing unused networks..."
docker network prune --force

echo "[APPOINTMENT] Removing unused images..."
docker image prune --force

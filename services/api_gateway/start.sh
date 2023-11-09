#!/bin/bash

echo "[IDM] Stopping existing containers..."
docker compose down

echo "[IDM] Building Docker images..."
docker compose build

echo "[IDM] Starting containers..."
docker compose up --force-recreate

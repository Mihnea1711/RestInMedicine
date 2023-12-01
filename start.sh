#!/bin/bash

echo "[DOCKER] Stopping and removing existing containers..."
docker-compose down --volumes --remove-orphans

echo "[DOCKER] Cleaning up images and containers..."
docker ps -aq --filter name="pos_project_*" | xargs docker stop | xargs docker rm
docker images --format="{{.Repository}}" | grep pos_project_ | xargs docker rmi

echo "[DOCKER] Pruning unused resources..."
docker container prune --force
docker image prune --force
docker volume prune --force

echo "[DOCKER] Building Docker images..."
docker compose build

echo "[DOCKER] Starting containers..."
docker compose up --force-recreate --build

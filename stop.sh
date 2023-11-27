#!/bin/bash

echo "[GATEWAY] Stopping and removing existing containers..."
docker-compose down --volumes --remove-orphans

# echo "[GATEWAY] Removing specific volumes..."
# docker volume rm pos_project_GATEWAY-mysql-data pos_project_GATEWAY-redis-data

echo "[GATEWAY] Cleaning up images and containers..."
docker ps -aq --filter name="pos_project_*" | xargs docker stop | xargs docker rm
docker images --format="{{.Repository}}" | grep pos_project_ | xargs docker rmi

echo "[GATEWAY] Pruning unused resources..."
docker container prune --force
docker image prune --force
docker volume prune --force
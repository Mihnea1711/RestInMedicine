#!/bin/bash

echo "[DOCTOR] Stopping existing containers..."
docker compose down

echo "[DOCTOR] Removing MySQL data volume..."
docker volume rm pos_project_doctori-mysql-data

echo "[DOCTOR] Removing Redis data volume..."
docker volume rm pos_project_doctori-redis-data

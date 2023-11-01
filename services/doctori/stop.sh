#!/bin/bash

echo "[DOCTOR] Stopping existing containers..."
docker compose down

echo "[DOCTOR] Removing MySQL data volume..."
docker volume rm doctori_mysql_data

echo "[DOCTOR] Removing Redis data volume..."
docker volume rm doctori_redis_data

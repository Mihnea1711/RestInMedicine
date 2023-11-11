#!/bin/bash

echo "[IDM] Stopping existing containers..."
docker compose down

echo "[IDM] Removing MySQL data volume..."
docker volume rm idm_mysql_data

echo "[IDM] Removing Redis data volume..."
docker volume rm idm_redis_data

echo "[IDM] Removing Images..."
docker rmi app_idm mysql:8.2.0 redis:latest

#!/bin/bash

# Stop and remove all containers:
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)

# Remove all images:
docker rmi $(docker images -a -q)

# Remove all volumes:
docker volume prune

# Remove all networks:
docker network prune

# Remove all caches:
docker builder prune

# Remove all system-wide unused resources:
docker system prune -a

# Remove all unused volumes, networks, and images:
docker system prune --volumes

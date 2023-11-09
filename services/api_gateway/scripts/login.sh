#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl \
    -X POST http://localhost:"$PORT"/api/login \
    -H "Content-Type: application/json" \
    -d '{
        "username": "mihnea",
        "password": "mihnea"
    }'

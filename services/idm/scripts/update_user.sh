#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl \
    -X PUT http://localhost:"$PORT"/idm/user/1 \
    -H "Content-Type: application/json" \
    -d '{
    "username": "new_username",
    "password": "new_password"
}' 

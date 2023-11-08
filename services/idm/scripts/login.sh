#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl -X POST -H "Content-Type: application/json" -d '{
    "username": "your_username",
    "password": "your_password",
    "role": "user_role"
}' http://localhost:8080/idm/register

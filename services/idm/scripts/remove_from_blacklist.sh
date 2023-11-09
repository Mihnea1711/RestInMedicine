#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# You should provide the necessary data in the request body
curl \
    -X POST http://localhost:"$PORT"/idm/blacklist/remove \
    -H "Content-Type: application/json" \
    -d '{
        "id_user": 1
    }'

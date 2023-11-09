#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Replace {id} with the actual user ID
# You should provide the new role in the request body
curl \
    -X PUT http://localhost:"$PORT"/idm/user/1/role \
    -H "Content-Type: application/json" \
    -d '{
        "role": "pacient" 
    }'



#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Replace {id} with the actual user ID
curl -X GET http://localhost:"$PORT"/idm/user/1/role
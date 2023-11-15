#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

USER_ID=1

# Use the extracted port in curl or other commands
curl -X GET http://localhost:"$PORT"/doctors/users/"$USER_ID"
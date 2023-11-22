#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Status you want to retrieve programari for 
STATUS="onorata"

# Use the extracted port in curl to retrieve programari by Status
curl -X GET http://localhost:"$PORT"/appointments/status/"$STATUS"

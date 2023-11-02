#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Programare ID you want to delete
PROGRAMARE_ID=1

# Use the extracted port in curl or other commands
curl -X DELETE http://localhost:"$PORT"/programari/"$PROGRAMARE_ID"

#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Programare ID you want to retrieve
CONSULTATIE_ID=000000000000000000000000

# Use the extracted port in curl or other commands
curl -X GET http://localhost:"$PORT"/consultations/"$CONSULTATIE_ID"

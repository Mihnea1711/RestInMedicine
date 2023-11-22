#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Pacient ID you want to retrieve appointments for
PACIENT_ID=1

# Use the extracted port in curl to retrieve appointments by Pacient ID
curl -X GET http://localhost:"$PORT"/appointments/pacient/"$PACIENT_ID"

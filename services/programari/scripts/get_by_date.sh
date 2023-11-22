#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Date you want to retrieve appointments for (replace with a valid date)
DATE="2023-11-10"

# Use the extracted port in curl to retrieve appointments by Date
curl -X GET http://localhost:"$PORT"/appointments/date/"$DATE"

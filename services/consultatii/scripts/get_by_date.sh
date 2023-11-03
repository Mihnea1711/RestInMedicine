#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Date you want to retrieve programari for (replace with a valid date)
DATE="2023-11-04"

# Use the extracted port in curl to retrieve programari by Date
curl -X GET http://localhost:"$PORT"/consultatii/date/"$DATE"

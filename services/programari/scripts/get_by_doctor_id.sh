#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Doctor ID you want to retrieve programari for
DOCTOR_ID=1

# Use the extracted port in curl to retrieve programari by Doctor ID
curl -X GET http://localhost:"$PORT"/programari/doctor/"$DOCTOR_ID"

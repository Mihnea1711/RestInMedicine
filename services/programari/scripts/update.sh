#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Programare ID you want to update
PROGRAMARE_ID=1

# Use the extracted port in curl or other commands
curl \
      -X PUT http://localhost:"$PORT"/programari/"$PROGRAMARE_ID" \
      -H "Content-Type: application/json"    \
      -d '{
          "idPacient": 2, 
          "idDoctor": 1, 
          "date": "2023-11-15T15:30:00Z", 
          "status": "anulata"
      }' 

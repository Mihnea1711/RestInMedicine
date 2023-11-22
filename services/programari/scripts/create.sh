#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Use the extracted port in curl or other commands
curl \
      -X POST http://localhost:"$PORT"/appointments \
      -H "Content-Type: application/json"    \
      -d '{
          "idPacient": 1, 
          "idDoctor": 1, 
          "date": "2023-10-31T14:00:00Z", 
          "status": "onorata"
      }' 


#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Use the extracted port in curl or other commands
curl -X PUT http://localhost:"$PORT"/doctori/1 \
     -H "Content-Type: application/json" \
     -d '{
        "idDoctor": 1, 
        "idUser": 123, 
        "nume": "Popescu", 
        "prenume": "Alexandru", 
        "email": "alex.popescu@example.com", 
        "telefon": "0743991354", 
        "specializare": "Neurologie"
        }'


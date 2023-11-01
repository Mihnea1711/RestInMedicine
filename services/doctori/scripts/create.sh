#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Use the extracted port in curl or other commands
curl \
      -X POST http://localhost:"$PORT"/doctori \
      -H "Content-Type: application/json"    \
      -d '{
          "idDoctor": 1, 
          "idUser": 123, 
          "nume": "Popescu", 
          "prenume": "Ion", 
          "email": "ion.popescu25@example.com", 
          "telefon": "0743991353", 
          "specializare": "Cardiologie"
      }' 

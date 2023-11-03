#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the JSON payload for creating a Consultatie
JSON_PAYLOAD='
{
  "id_pacient": 1,
  "id_doctor": 2,
  "data": "2023-11-03T00:00:00Z",
  "diagnostic": "Sample diagnostic",
  "investigatii": [
    {
      "denumire": "Investigatie 1",
      "durata_de_procesare": 60,
      "rezultat": "Rezultat 1"
    }
  ]
}'

# Use curl to create a new Consultatie
curl \
    -X POST http://localhost:"$PORT"/consultatii \
    -H "Content-Type: application/json" \
    -d "$JSON_PAYLOAD" 


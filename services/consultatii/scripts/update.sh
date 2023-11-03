#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Specify the Consultatie ID you want to update
CONSULTATIE_ID=000000000000000000000000

# Specify the JSON payload for updating a Consultatie
JSON_PAYLOAD='
{
  "id_pacient": 1,
  "id_doctor": 2,
  "data": "2023-11-03T14:30:00Z",
  "diagnostic": "Updated diagnostic",
  "investigatii": [
    {
      "denumire": "Updated Investigatie 1",
      "durata_de_procesare": 90,
      "rezultat": "Updated Rezultat 1"
    }
  ]
}'

# Use curl to update a Consultatie by ID
curl \
    -X PUT http://localhost:"$PORT"/consultatii/"$CONSULTATIE_ID" \
    -H "Content-Type: application/json" \
    -d "$JSON_PAYLOAD" 

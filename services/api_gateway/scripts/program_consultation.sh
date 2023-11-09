#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl \
    -X POST http://localhost:"$PORT"/api/consultations \
    -H "Content-Type: application/json" \
    -d '{
        "id_patient": 789,
        "id_doctor": 101,
        "date": "2023-12-15T15:30:00"
    }'

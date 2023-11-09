#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl \
    -X POST http://localhost:"$PORT"/api/appointments \
    -H "Content-Type: application/json" \
    -d '{
        "id_patient": 123,
        "id_doctor": 456,
        "date": "2023-12-01T10:00:00"
    }'

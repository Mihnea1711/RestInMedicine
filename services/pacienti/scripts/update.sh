#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl \
    -X PUT http://localhost:"$PORT"/patients/1 \
    -H "Content-Type: application/json" \
    -d '{
        "id_user": 1,
        "nume": "UpdatedName",
        "prenume": "UpdatedPrenume",
        "email": "updated@example.com",
        "telefon": "0712345678",
        "cnp": "1000101890123",
        "data_nasterii": "2000-01-01T00:00:00Z",
        "is_active": true
    }'


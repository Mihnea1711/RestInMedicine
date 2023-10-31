#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl \
    -X PUT http://localhost:"$PORT"/pacienti/1 \
    -H "Content-Type: application/json" \
    -d '{
        "cnp": "1234567890123",
        "id_user": 1,
        "nume": "UpdatedName",
        "prenume": "UpdatedPrenume",
        "email": "updated@example.com",
        "telefon": "0712345678",
        "data_nasterii": "2000-01-01",
        "is_active": true
    }'

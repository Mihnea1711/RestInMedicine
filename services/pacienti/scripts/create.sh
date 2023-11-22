#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl \
    -X POST http://localhost:"$PORT"/patients \
    -H "Content-Type: application/json" \
    -d '{
        "id_user": 5,
        "nume": "John",
        "prenume": "Doe",
        "email": "johndoe@example.com",
        "telefon": "0712345678",
        "cnp": "1000101190123",
        "data_nasterii": "2000-01-01T00:00:00Z",
        "is_active": true
    }'


#!/bin/bash
# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

ROLE="admin"

echo http://localhost:"$PORT"/api/register?role="$ROLE"
curl -X POST http://localhost:"$PORT"/api/register?role="$ROLE"
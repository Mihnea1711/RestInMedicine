#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl -i -w "\n" http://localhost:"$PORT"/api/rabbit/health-check

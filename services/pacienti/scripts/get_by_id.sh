#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

# Use the extracted port in curl or other commands
curl -X GET http://localhost:"$PORT"/pacienti/1
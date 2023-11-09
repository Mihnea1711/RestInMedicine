#!/bin/bash

# Extract port from config.yaml
PORT=$(yq e '.server.port' configs/config.yaml)

curl -X DELETE http://localhost:"$PORT"/idm/user/1

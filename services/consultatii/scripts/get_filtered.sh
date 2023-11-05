#!/bin/bash

# Set the port and endpoint
FILTER_CONSULTATII_ENDPOINT="/consultatii/filter"
PORT=$(yq e '.server.port' configs/config.yaml)

# Define query parameters for the filter (you can modify these as needed)
# DATE="2023-11-17"
DATE="\"2023-11-17\""
PACIENT_ID=3
DOCTOR_ID=3

echo "http://localhost:$PORT$FILTER_CONSULTATII_ENDPOINT?date=$DATE&id_pacient=$PACIENT_ID&id_doctor=$DOCTOR_ID"

# Make a GET request to the filter endpoint with query parameters
curl -X GET "http://localhost:$PORT$FILTER_CONSULTATII_ENDPOINT?date=$DATE&id_pacient=$PACIENT_ID&id_doctor=$DOCTOR_ID"
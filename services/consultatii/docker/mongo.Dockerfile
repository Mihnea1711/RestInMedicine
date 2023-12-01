# Use the MongoDB base image
FROM mongo:latest

# Copy the initialization script to the container
# COPY ./docker/db/init-mongodb.js /data/db 
COPY ./docker/db/init-mongodb.js /docker-entrypoint-initdb.d/

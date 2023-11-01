    # Use Go 1.21 as the base image
    FROM mysql:8.2.0

    COPY ./docker/db/*.sql /docker-entrypoint-initdb.d/

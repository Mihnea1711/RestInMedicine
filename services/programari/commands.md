# App Programari User Guide

## Building the Module Locally
To build the module, run the following command:

```bash
go build -o app_programari
```

## Running the Module Locally
To run the module, run the following command:

```bash
./app_programari
```

## Testing the Module

### Testing Create Route
To test the create programare route or to change request payload, run the following command:

```bash
chmod +x scripts/create.sh
./scripts/create.sh
```

### Testing Get All Route
To test the get all programari route, run the following command:

```bash
chmod +x scripts/get_all.sh
./scripts/get_all.sh

```

### Testing get By ID Route
To test the get programare by id route or to change request payload, go to /scripts/get_by_id.sh, change contents and run the following command:

```bash
chmod +x scripts/get_by_id.sh
./scripts/get_by_id.sh
```

### Testing Update Route
To test the update programare by id route or to change request payload, change contents and run the following command:

```bash
chmod +x scripts/update.sh
./scripts/update.sh
```

### Testing Delete Route
To test the delete programare by id route or to change request payload, change contents and run the following command:

```bash
chmod +x scripts/delete.sh
./scripts/delete.sh
```

## MySQL Info
### Pull the image locally with:

```bash
docker pull mysql:8.2.0
```

### Run the image locally with:

```bash
docker run --network="host" --name mysql-container -e MYSQL_ROOT_PASSWORD=mihnea_pos -d -p 3306:3306 mysql:8.2.0
```
or
```bash
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=mihnea_pos -d -p 3306:3306 mysql:8.2.0
```

### Stop the image with:

```bash
docker stop mysql-container     # or docker stop <CONTAINER_ID>
```

### Exec the mysql (Get into mysql cli) with:

```bash
docker exec -it mysql-container mysql -u root -p
docker exec -it mysql-container mysql -u mihnea_pos -p
```

or

```bash
docker exec -it <Container_DB_Id> /bin/bash         # get container id from docker ps command
mysql -u mihnea_pos -pmihnea_pos pos_db

```

### Useful commands
CREATE DATABASE your_database_name;
SHOW DATABASES;
USE your_database_name;
SHOW TABLES;
DESCRIBE your_table_name;
SELECT * FROM your_table_name;
EXIT;

CREATE USER 'username'@'host' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON your_database_name.* TO 'username'@'host';
FLUSH PRIVILEGES;

## Deploy with Docker Compose

```bash
docker compose down

docker volume rm programari_mysql_data
docker volume rm programari_redis_data

docker compose build 
docker compose up --force-recreate
```
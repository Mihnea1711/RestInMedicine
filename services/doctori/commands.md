# App Doctori User Guide
PORT = 8080

## Building the Module
To build the module, run the following command:

```bash
go build -o app_doctori
```

## Running the Module
To run the module, run the following command:

```bash
./app_doctori
```

## Testing the Module

### Testing Create Route
To test the create doctor route or to change request payload, run the following command:

```bash
chmod +x scripts/create.sh
./scripts/create.sh
```

### Testing Get All Route
To test the get all doctors route, run the following command:

```bash
chmod +x scripts/get_all.sh
./scripts/get_all.sh

```

### Testing get By ID Route
To test the get doctor by id route or to change request payload, go to /scripts/get_by_id.sh, change contents and run the following command:

```bash
chmod +x scripts/get_by_id.sh
./scripts/get_by_id.sh
```

### Testing Update Route
To test the update doctor by id route or to change request payload, change contents and run the following command:

```bash
chmod +x scripts/update.sh
./scripts/update.sh
```

### Testing Delete Route
To test the delete doctor by id route or to change request payload, change contents and run the following command:

```bash
chmod +x scripts/delete.sh
./scripts/delete.sh
```
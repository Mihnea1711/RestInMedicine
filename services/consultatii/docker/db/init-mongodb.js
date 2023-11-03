// Define your database name
var dbName = "mydatabase";

// Connect to the database
var db = db.getSiblingDB(dbName);

// Create a collection and insert data
db.programare.insertMany([
    {
        id_pacient: 1,
        id_doctor: 1,
        date: new ISODate("2023-11-10T00:00:00Z"),
        status: "onorata"
    },
    {
        id_pacient: 2,
        id_doctor: 2,
        date: new ISODate("2023-11-15T00:00:00Z"),
        status: "anulata"
    }
]);

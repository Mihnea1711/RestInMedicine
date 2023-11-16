// Define your database name
var dbName = "consultations_db";

// Connect to the database
var db = db.getSiblingDB(dbName);

// Define the collection name SAME AS THE CONTSTANTS
var collectionName = "consultation";

// Define two Consultatie objects
var consultatie1 = {
    id_pacient: 1,
    id_doctor: 1,
    date: new ISODate("2023-11-10"),
    diagnostic: "Diagnostic for Consultatie 1",
    investigatii: [
        {
            denumire: "Investigatia 1",
            durata_procesare: 30,
            rezultat: "Rezultat investigatie 1"
        }
    ]
};

var consultatie2 = {
    id_pacient: 2,
    id_doctor: 2,
    date: new ISODate("2023-11-15"),
    diagnostic: "Diagnostic for Consultatie 2",
    investigatii: [
        {
            denumire: "Investigatia 2",
            durata_procesare: 45,
            rezultat: "Rezultat investigatie 2"
        }
    ]
};

// Insert the Consultatie objects into the collection
db[collectionName].insertMany([consultatie1, consultatie2]);
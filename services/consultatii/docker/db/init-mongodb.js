// Define your database name
var dbName = "consultations_db";

// Connect to the admin database
var adminDB = db.getSiblingDB("admin");

// Create a user with the necessary privileges for the targeted database
adminDB.createUser({
  user: "mihnea_pos",
  pwd: "mihnea_pos",
  roles: [
    { role: "readWrite", db: dbName },
    { role: "dbAdmin", db: dbName }
  ]
});

// Connect to the database
var db = db.getSiblingDB(dbName);

// Define the collection name SAME AS THE CONTSTANTS
var collectionName = "consultation";

db[collectionName].createIndex({ id_patient: 1, id_doctor: 1, date: 1 }, { unique: true });

// Define two Consultatie objects
var consultatie1 = {
    id_consultation: ObjectId(),
    id_patient: 1,
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
    id_consultation: ObjectId(),
    id_patient: 2,
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
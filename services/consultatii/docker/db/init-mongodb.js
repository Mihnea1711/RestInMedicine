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

// Define two Consultation objects
var consultation1 = {
    _id: ObjectId(),
    id_patient: 1,
    id_doctor: 1,
    date: new ISODate("2023-11-10"),
    diagnostic: "Diagnostic for Consultation 1",
    investigations: [
        {
            id_investigation: ObjectId(),
            name: "Investigation 1",
            processingTime: 30,
            result: "Result of Investigation 1"
        }
    ]
};

var consultation2 = {
    _id: ObjectId(),
    id_patient: 2,
    id_doctor: 2,
    date: new ISODate("2023-11-15"),
    diagnostic: "Diagnostic for Consultation 2",
    investigations: [
        {
            id_investigation: ObjectId(),
            name: "Investigation 2",
            processingTime: 45,
            result: "Result of Investigation 2"
        }
    ]
};

// Insert the Consultatie objects into the collection
db[collectionName].insertMany([consultation1, consultation2]);
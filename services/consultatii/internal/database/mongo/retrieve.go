package mongo

import (
	"context"
	"log"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchAllConsultations retrieves all consultations.
func (db *MongoDB) FetchAllConsultations(ctx context.Context, page, limit int) ([]models.Consultation, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Set options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	// Perform the find operation
	cur, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error fetching all consultations: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	// Slice to store retrieved consultations
	var consultations []models.Consultation

	// Iterate through the result cursor
	for cur.Next(ctx) {
		var consultation models.Consultation
		// Decode each document into a consultation struct
		if err := cur.Decode(&consultation); err != nil {
			log.Printf("[CONSULTATION] Error decoding consultation: %v", err)
			return nil, err
		}
		consultations = append(consultations, consultation)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATION] Error iterating over results: %v", err)
		return nil, err
	}

	// Log results
	if len(consultations) == 0 {
		log.Println("[CONSULTATION] No consultations found.")
	} else {
		log.Printf("[CONSULTATION] Retrieved %d consultations successfully.", len(consultations))
	}

	return consultations, nil
}

// FetchConsultationByID retrieves a consultation by its ObjectID.
func (db *MongoDB) FetchConsultationByID(ctx context.Context, consultationID primitive.ObjectID) (*models.Consultation, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Log the ID being fetched
	log.Printf("[CONSULTATION] Attempting to fetch consultation by ID: %s", consultationID.String())

	// Create a models.Consultation instance to store the result
	var consultation models.Consultation

	// Attempt to find the consultation by ID
	err := collection.FindOne(ctx, bson.M{utils.ID_CONSULTATIE: consultationID}).Decode(&consultation)
	if err != nil {
		// Handle other errors
		log.Printf("[CONSULTATION] Error fetching consultation: %v", err)
		return nil, err
	}

	// Log the successful retrieval of the consultation
	log.Printf("[CONSULTATION] Consultation retrieved successfully: %v", consultationID)
	return &consultation, nil
}

// FetchConsultationsByPatientID retrieves consultations by PacientID.
func (db *MongoDB) FetchConsultationsByPatientID(ctx context.Context, patientID int, page, limit int) ([]models.Consultation, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Define the filter to find consultations by PacientID
	filter := bson.M{utils.ID_PACIENT: patientID}

	// Configure options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	// Execute the query and get a cursor to the results
	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error fetching consultations by PacientID: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	// Initialize a slice to store the fetched consultations
	var consultations []models.Consultation

	// Iterate over the cursor and decode each consultation
	for cur.Next(ctx) {
		var consultation models.Consultation
		if err := cur.Decode(&consultation); err != nil {
			log.Printf("[CONSULTATION] Error decoding consultation: %v", err)
			return nil, err
		}
		consultations = append(consultations, consultation)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATION] Error iterating over results: %v", err)
		return nil, err
	}

	// Log the successful retrieval of consultations by PacientID
	log.Printf("[CONSULTATION] Consultations by PacientID retrieved successfully: %v", patientID)
	return consultations, nil
}

// FetchConsultationsByDoctorID retrieves consultations by DoctorID.
func (db *MongoDB) FetchConsultationsByDoctorID(ctx context.Context, doctorID int, page, limit int) ([]models.Consultation, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Define the filter to find consultations by DoctorID
	filter := bson.M{utils.ID_DOCTOR: doctorID}

	// Configure options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	// Execute the query and get a cursor to the results
	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error fetching consultations by DoctorID: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	// Initialize a slice to store the fetched consultations
	var consultations []models.Consultation

	// Iterate over the cursor and decode each consultation
	for cur.Next(ctx) {
		var consultation models.Consultation
		if err := cur.Decode(&consultation); err != nil {
			log.Printf("[CONSULTATION] Error decoding consultation: %v", err)
			return nil, err
		}
		consultations = append(consultations, consultation)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATION] Error iterating over results: %v", err)
		return nil, err
	}

	// Log the successful retrieval of consultations by DoctorID
	log.Printf("[CONSULTATION] Consultations by DoctorID retrieved successfully: %v", doctorID)
	return consultations, nil
}

// FetchConsultationsByDate retrieves consultations by Date.
func (db *MongoDB) FetchConsultationsByDate(ctx context.Context, date time.Time, page, limit int) ([]models.Consultation, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Define the filter to find consultations by Date
	filter := bson.M{utils.DATE: bson.M{"$eq": date}}

	// Configure options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	// Execute the query and get a cursor to the results
	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error fetching consultations by Date: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	// Initialize a slice to store the fetched consultations
	var consultations []models.Consultation

	// Iterate over the cursor and decode each consultation
	for cur.Next(ctx) {
		var consultation models.Consultation
		if err := cur.Decode(&consultation); err != nil {
			log.Printf("[CONSULTATION] Error decoding consultation: %v", err)
			return nil, err
		}
		consultations = append(consultations, consultation)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATION] Error iterating over results: %v", err)
		return nil, err
	}

	// Log the successful retrieval of consultations by Date
	log.Printf("[CONSULTATION] Consultations by Date retrieved successfully: %v", date)
	return consultations, nil
}

// FetchConsultationsByFilter retrieves consultations based on the provconsultationIDed filter criteria.
func (db *MongoDB) FetchConsultationsByFilter(ctx context.Context, filter bson.D, page int, limit int) ([]models.Consultation, error) {
	// Create a MongoDB cursor for querying the collection
	collection := db.client.Database(utils.DATABASE_NAME).Collection(utils.CONSULTATIE_TABLE)

	// log.Println(collection.Name())
	// log.Println(collection.Database().Name())

	// timeV, _ := time.Parse(utils.TIME_FORMAT, "2023-11-17")

	custom_f := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "consultationID_doctor", Value: 2}},
				bson.D{{Key: "consultationID_pacient", Value: 2}},
				// bson.D{{Key: "date", Value: timeV}},
			}},
	}

	cfilter := bson.D{
		{Key: "consultationID_doctor", Value: 2},
		{Key: "consultationID_pacient", Value: 2},
		// {Key: "date", Value: timeV},
	}

	log.Printf("Filter: %v", filter)
	log.Printf("Custom_F: %v", custom_f)
	log.Printf("C_Filter: %v", cfilter)

	// Define options for pagination
	options := options.Find()
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	// -----------------------------------

	// Log the filter parameters for debugging
	log.Printf("[CONSULTATION] Fetching consultations with filter: %v, Limit of %d, on Page %d", filter, limit, page)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("[CONSULTATION] Failed to find consultations with filter: %v", filter)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Fetch consultations based on the filter and pagination options
	var consultations []models.Consultation
	if err := cursor.All(ctx, &consultations); err != nil {
		log.Printf("[CONSULTATION] Failed to decode consultations: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Fetched %d consultations with filter: %v", len(consultations), filter)

	log.Println("-----------------------------")

	// -----------------------------------

	// Log the filter parameters for debugging
	log.Printf("[CONSULTATION] Fetching consultations with filter: %v, Limit of %d, on Page %d", custom_f, limit, page)

	cursor, err = collection.Find(ctx, custom_f)
	if err != nil {
		log.Printf("[CONSULTATION] Failed to find consultations with filter: %v", custom_f)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Fetch consultations based on the filter and pagination options
	var consultations2 []models.Consultation
	if err := cursor.All(ctx, &consultations2); err != nil {
		log.Printf("[CONSULTATION] Failed to decode consultations: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Fetched %d consultations with filter: %v", len(consultations), custom_f)

	log.Println("-----------------------------")

	// -----------------------------------

	// Log the filter parameters for debugging
	log.Printf("[CONSULTATION] Fetching consultations with filter: %v, Limit of %d, on Page %d", cfilter, limit, page)

	cursor, err = collection.Find(ctx, cfilter)
	if err != nil {
		log.Printf("[CONSULTATION] Failed to find consultations with filter: %v", cfilter)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Fetch consultations based on the filter and pagination options
	var consultations3 []models.Consultation
	if err := cursor.All(ctx, &consultations3); err != nil {
		log.Printf("[CONSULTATION] Failed to decode consultations: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Fetched %d consultations with filter: %v", len(consultations), cfilter)

	log.Println("-----------------------------")

	// -----------------------------------

	log.Println(consultations)
	log.Println(consultations2)
	log.Println(consultations3)

	return nil, nil
}

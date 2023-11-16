package mongo

import (
	"context"
	"log"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchAllConsultatii retrieves all consultatii.
func (db *MongoDB) FetchAllConsultatii(ctx context.Context, page, limit int) ([]models.Consultatie, error) {
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	cur, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error fetching all consultatii: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var consultatii []models.Consultatie
	for cur.Next(ctx) {
		var consultatie models.Consultatie
		if err := cur.Decode(&consultatie); err != nil {
			log.Printf("[CONSULTATION] Error decoding consultatie: %v", err)
			return nil, err
		}
		consultatii = append(consultatii, consultatie)
	}

	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATION] Error iterating over results: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] All consultatii retrieved successfully.")
	return consultatii, nil
}

// FetchConsultatieByID retrieves a consultatie by its ObjectID.
func (db *MongoDB) FetchConsultatieByID(ctx context.Context, id primitive.ObjectID) (*models.Consultatie, error) {
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	var consultatie models.Consultatie
	err := collection.FindOne(ctx, bson.M{utils.ID_CONSULTATIE: id}).Decode(&consultatie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle the case where no document is found with the given ID
			log.Printf("[CONSULTATION] Consultatie not found: %v", id)
			return nil, nil
		}

		// Handle other errors
		log.Printf("[CONSULTATION] Error fetching consultatie: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Consultatie retrieved successfully: %v", id)
	return &consultatie, nil
}

// FetchConsultatiiByPacientID retrieves consultatii by PacientID.
func (db *MongoDB) FetchConsultatiiByPacientID(ctx context.Context, pacientID int, page, limit int) ([]models.Consultatie, error) {
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)
	filter := bson.M{utils.ID_PACIENT: pacientID}

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error fetching consultatii by PacientID: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var consultatii []models.Consultatie
	for cur.Next(ctx) {
		var consultatie models.Consultatie
		if err := cur.Decode(&consultatie); err != nil {
			log.Printf("[CONSULTATION] Error decoding consultatie: %v", err)
			return nil, err
		}
		consultatii = append(consultatii, consultatie)
	}

	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATION] Error iterating over results: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Consultatii by PacientID retrieved successfully: %v", pacientID)
	return consultatii, nil
}

// FetchConsultatiiByDoctorID retrieves consultatii by DoctorID.
func (db *MongoDB) FetchConsultatiiByDoctorID(ctx context.Context, doctorID int, page, limit int) ([]models.Consultatie, error) {
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)
	filter := bson.M{utils.ID_DOCTOR: doctorID}

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	log.Printf("LIMIT / PAGE: %d, %d", limit, page)
	log.Printf("RETRIEVE: %d", doctorID)

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error fetching consultatii by DoctorID: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var consultatii []models.Consultatie
	for cur.Next(ctx) {
		var consultatie models.Consultatie
		if err := cur.Decode(&consultatie); err != nil {
			log.Printf("[CONSULTATION] Error decoding consultatie: %v", err)
			return nil, err
		}
		consultatii = append(consultatii, consultatie)
	}

	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATION] Error iterating over results: %v", err)
		return nil, err
	}

	log.Printf("consultatii: %v", consultatii)

	log.Printf("[CONSULTATION] Consultatii by DoctorID retrieved successfully: %v", doctorID)
	return consultatii, nil
}

// FetchConsultatiiByDate retrieves consultatii by Date.
func (db *MongoDB) FetchConsultatiiByDate(ctx context.Context, date time.Time, page, limit int) ([]models.Consultatie, error) {
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)
	filter := bson.M{utils.DATE: bson.M{"$eq": date}}

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error fetching consultatii by Date: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var consultatii []models.Consultatie
	for cur.Next(ctx) {
		var consultatie models.Consultatie
		if err := cur.Decode(&consultatie); err != nil {
			log.Printf("[CONSULTATION] Error decoding consultatie: %v", err)
			return nil, err
		}
		consultatii = append(consultatii, consultatie)
	}

	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATION] Error iterating over results: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Consultatii by Date retrieved successfully: %v", date)
	return consultatii, nil
}

// FetchConsultatiiByFilter retrieves consultatii based on the provided filter criteria.
func (db *MongoDB) FetchConsultatiiByFilter(ctx context.Context, filter bson.D, page int, limit int) ([]models.Consultatie, error) {
	// Create a MongoDB cursor for querying the collection
	collection := db.client.Database(utils.DATABASE_NAME).Collection(utils.CONSULTATIE_TABLE)

	// log.Println(collection.Name())
	// log.Println(collection.Database().Name())

	// timeV, _ := time.Parse(utils.TIME_FORMAT, "2023-11-17")

	custom_f := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "id_doctor", Value: 2}},
				bson.D{{Key: "id_pacient", Value: 2}},
				// bson.D{{Key: "date", Value: timeV}},
			}},
	}

	cfilter := bson.D{
		{Key: "id_doctor", Value: 2},
		{Key: "id_pacient", Value: 2},
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
	log.Printf("[CONSULTATION] Fetching consultatii with filter: %v, Limit of %d, on Page %d", filter, limit, page)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("[CONSULTATION] Failed to find consultatii with filter: %v", filter)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Fetch consultatii based on the filter and pagination options
	var consultatii []models.Consultatie
	if err := cursor.All(ctx, &consultatii); err != nil {
		log.Printf("[CONSULTATION] Failed to decode consultatii: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Fetched %d consultatii with filter: %v", len(consultatii), filter)

	log.Println("-----------------------------")

	// -----------------------------------

	// Log the filter parameters for debugging
	log.Printf("[CONSULTATION] Fetching consultatii with filter: %v, Limit of %d, on Page %d", custom_f, limit, page)

	cursor, err = collection.Find(ctx, custom_f)
	if err != nil {
		log.Printf("[CONSULTATION] Failed to find consultatii with filter: %v", custom_f)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Fetch consultatii based on the filter and pagination options
	var consultatii2 []models.Consultatie
	if err := cursor.All(ctx, &consultatii2); err != nil {
		log.Printf("[CONSULTATION] Failed to decode consultatii: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Fetched %d consultatii with filter: %v", len(consultatii), custom_f)

	log.Println("-----------------------------")

	// -----------------------------------

	// Log the filter parameters for debugging
	log.Printf("[CONSULTATION] Fetching consultatii with filter: %v, Limit of %d, on Page %d", cfilter, limit, page)

	cursor, err = collection.Find(ctx, cfilter)
	if err != nil {
		log.Printf("[CONSULTATION] Failed to find consultatii with filter: %v", cfilter)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Fetch consultatii based on the filter and pagination options
	var consultatii3 []models.Consultatie
	if err := cursor.All(ctx, &consultatii3); err != nil {
		log.Printf("[CONSULTATION] Failed to decode consultatii: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATION] Fetched %d consultatii with filter: %v", len(consultatii), cfilter)

	log.Println("-----------------------------")

	// -----------------------------------

	log.Println(consultatii)
	log.Println(consultatii2)
	log.Println(consultatii3)

	return nil, nil
}

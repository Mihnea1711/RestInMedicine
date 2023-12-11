package mongo

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchConsultationsByFilter retrieves consultations based on the provconsultationIDed filter criteria.
func (db *MongoDB) FetchConsultations(ctx context.Context, filter bson.M, page int, limit int) ([]models.Consultation, error) {
	// Create a MongoDB cursor for querying the collection
	collection := db.client.Database(utils.DATABASE_NAME).Collection(utils.CONSULTATIE_TABLE)

	// Define options for pagination
	options := options.Find()
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	// Log the filter parameters for debugging
	log.Printf("[CONSULTATION] Fetching consultations with filter: %v, Limit of %d, on Page %d", filter, limit, page)

	cursor, err := collection.Find(ctx, filter, options)
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

	// Log results
	if len(consultations) == 0 {
		log.Println("[CONSULTATION] No consultations found in database.")
	} else {
		log.Printf("[CONSULTATION] Fetched %d consultations", len(consultations))
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
	err := collection.FindOne(ctx, bson.M{utils.COLUMN_ID_CONSULTATIE: consultationID}).Decode(&consultation)
	if err != nil {
		// Handle other errors
		log.Printf("[CONSULTATION] Error fetching consultation: %v", err)
		return nil, err
	}

	// Log the successful retrieval of the consultation
	log.Printf("[CONSULTATION] Consultation retrieved successfully: %v", consultationID)
	return &consultation, nil
}

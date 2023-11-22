package mongo

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateConsultationByID updates a consultation by its ObjectID and returns the number of rows affected.
func (db *MongoDB) UpdateConsultationByID(ctx context.Context, consultation *models.Consultation) (int, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Log the attempt to update the consultation with its ID
	log.Printf("[PATIENT] Attempting to update consultation %s", consultation.IDConsultatie.Hex())

	// Replace the existing consultation document in the collection with the provided one
	result, err := collection.ReplaceOne(ctx, bson.M{utils.ID_CONSULTATIE: consultation.IDConsultatie}, consultation)
	if err != nil {
		// Log an error if the update operation fails
		log.Printf("[CONSULTATION] Error updating consultation by ID: %v", err)
		return 0, err
	}

	// Check if any document was modified during the update
	if result.ModifiedCount != 0 {
		// Log a success message if the consultation was updated
		log.Printf("[CONSULTATION] Consultation updated successfully. ID: %v", consultation.IDConsultatie.Hex())
	} else {
		// Log a message if no consultation was updated (ID not found)
		log.Printf("[CONSULTATION] No consultation has been updated for ID: %v", consultation.IDConsultatie.Hex())
	}

	// Return the number of modified documents (rows) and a potential error
	return int(result.ModifiedCount), nil
}

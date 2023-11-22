package mongo

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteConsultationByID deletes a consultation by its ObjectID and returns the number of rows affected.
func (db *MongoDB) DeleteConsultationByID(ctx context.Context, consultationID primitive.ObjectID) (int, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Perform the deletion operation based on the provided consultation ID
	result, err := collection.DeleteOne(ctx, bson.M{utils.ID_CONSULTATIE: consultationID})
	if err != nil {
		// Log an error if the deletion operation fails
		log.Printf("[CONSULTATION] Error deleting consultation by ID %v: %v", consultationID.Hex(), err)
		return 0, err
	}

	// Check if any document was deleted during the operation
	if result.DeletedCount != 0 {
		// Log a success message if the consultation was deleted
		log.Printf("[CONSULTATION] Consultation with ID %v deleted successfully.", consultationID.Hex())
	} else {
		// Log a message if no consultation was deleted (ID not found)
		log.Printf("[CONSULTATION] No consultation with ID %v has been deleted.", consultationID.Hex())
	}

	// Return the number of deleted documents (rows) and a potential error
	return int(result.DeletedCount), nil
}

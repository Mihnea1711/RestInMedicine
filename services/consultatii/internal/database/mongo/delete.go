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

// DeleteConsultationsByPatientID deletes a consultation by its patient or doctor ID and returns the number of rows affected.
func (db *MongoDB) DeleteConsultationsByPatientOrDoctorID(ctx context.Context, id int) (int, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Define the filter to find consultations associated with the given user ID
	filter := bson.M{
		"$or": []bson.M{
			{"id_patient": id},
			{"id_doctor": id},
		},
	}

	// Perform the deletion operation based on the provided consultation ID
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		// Log an error if the deletion operation fails
		log.Printf("[CONSULTATION] Error deleting consultations by patient or doctor ID %d: %v", id, err)
		return 0, err
	}

	// Check if any document was deleted during the operation
	if result.DeletedCount != 0 {
		// Log a success message if the consultation was deleted
		log.Printf("[CONSULTATION] Consultations with patient or doctor ID %d deleted successfully.", id)
	} else {
		// Log a message if no consultation was deleted (ID not found)
		log.Printf("[CONSULTATION] No consultation with patient or doctor ID %d has been deleted.", id)
	}

	// Return the number of deleted documents (rows) and a potential error
	return int(result.DeletedCount), nil
}

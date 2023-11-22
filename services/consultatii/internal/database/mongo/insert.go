package mongo

import (
	"context"
	"errors"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveConsultatie saves a consultation to the database.
func (db *MongoDB) SaveConsultation(ctx context.Context, consultation *models.Consultation) (primitive.ObjectID, error) {
	// Get the MongoDB collection for consultations
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	// Log the attempt to save the consultation
	log.Println("[PATIENT] Attempting to save consultation")

	// Insert the consultation document into the collection
	result, err := collection.InsertOne(ctx, consultation)
	if err != nil {
		// Log an error if the insertion operation fails
		log.Printf("[CONSULTATION] Error saving consultation: %v", err)
		return primitive.NilObjectID, err
	}

	// Check if the InsertedID is valid and convert it to primitive.ObjectID
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok || oid.IsZero() {
		// Log an error if obtaining the last insert ID fails
		log.Printf("[CONSULTATION] Error getting last insert ID when saving consultation.")
		return primitive.NilObjectID, errors.New("error getting last insert ID")
	}

	// Log a success message with the ID of the saved consultation
	log.Printf("[CONSULTATION] Consultatie saved successfully. ID: %v", oid.Hex())

	// Return the ObjectID of the saved consultation and a potential error
	return oid, nil
}

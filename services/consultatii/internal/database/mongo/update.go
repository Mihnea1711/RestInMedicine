package mongo

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateConsultatieByID updates a consultatie by its ObjectID and returns the number of rows affected.
func (db *MongoDB) UpdateConsultatieByID(ctx context.Context, consultatie *models.Consultatie) (int64, error) {
	collection := db.db.Collection("consultatii")

	result, err := collection.ReplaceOne(ctx, bson.M{"_id": consultatie.IDConsultatie}, consultatie)
	if err != nil {
		log.Printf("[CONSULTATIE] Error updating consultatie by ID: %v", err)
		return 0, err
	}

	log.Printf("[CONSULTATIE] Consultatie updated successfully.")
	return result.ModifiedCount, nil
}

package mongo

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateConsultatieByID updates a consultatie by its ObjectID and returns the number of rows affected.
func (db *MongoDB) UpdateConsultatieByID(ctx context.Context, consultatie *models.Consultatie) (int64, error) {
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	result, err := collection.ReplaceOne(ctx, bson.M{utils.ID_CONSULTATIE: consultatie.IDConsultatie}, consultatie)
	if err != nil {
		log.Printf("[CONSULTATION] Error updating consultatie by ID: %v", err)
		return 0, err
	}

	if result.ModifiedCount != 0 {
		log.Printf("[CONSULTATION] Consultatie updated successfully.")
	} else {
		log.Printf("[CONSULTATION] No consultatie has been updated.")
	}
	return result.ModifiedCount, nil
}

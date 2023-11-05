package mongo

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteConsultatieByID deletes a consultatie by its ObjectID and returns the number of rows affected.
func (db *MongoDB) DeleteConsultatieByID(ctx context.Context, id primitive.ObjectID) (int64, error) {
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	result, err := collection.DeleteOne(ctx, bson.M{utils.ID_CONSULTATIE: id})
	if err != nil {
		log.Printf("[CONSULTATIE] Error deleting consultatie by ID: %v", err)
		return 0, err
	}

	log.Printf("[CONSULTATIE] Consultatie deleted successfully.")
	return result.DeletedCount, nil
}

package mongo

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

// SaveConsultatie saves a consultatie to the database.
func (db *MongoDB) SaveConsultatie(ctx context.Context, consultatie *models.Consultatie) error {
	collection := db.db.Collection(utils.CONSULTATIE_TABLE)

	_, err := collection.InsertOne(ctx, consultatie)
	if err != nil {
		log.Printf("[CONSULTATIE] Error saving consultatie: %v", err)
		return err
	}

	log.Printf("[CONSULTATIE] Consultatie saved successfully.")
	return nil
}

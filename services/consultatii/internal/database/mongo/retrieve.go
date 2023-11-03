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
		log.Printf("[CONSULTATIE] Error fetching all consultatii: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var consultatii []models.Consultatie
	for cur.Next(ctx) {
		var consultatie models.Consultatie
		if err := cur.Decode(&consultatie); err != nil {
			log.Printf("[CONSULTATIE] Error decoding consultatie: %v", err)
			return nil, err
		}
		consultatii = append(consultatii, consultatie)
	}

	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATIE] Error iterating over results: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATIE] All consultatii retrieved successfully.")
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
			log.Printf("[CONSULTATIE] Consultatie not found: %v", id)
			return nil, nil
		}

		// Handle other errors
		log.Printf("[CONSULTATIE] Error fetching consultatie: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATIE] Consultatie retrieved successfully: %v", id)
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
		log.Printf("[CONSULTATIE] Error fetching consultatii by PacientID: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var consultatii []models.Consultatie
	for cur.Next(ctx) {
		var consultatie models.Consultatie
		if err := cur.Decode(&consultatie); err != nil {
			log.Printf("[CONSULTATIE] Error decoding consultatie: %v", err)
			return nil, err
		}
		consultatii = append(consultatii, consultatie)
	}

	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATIE] Error iterating over results: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATIE] Consultatii by PacientID retrieved successfully: %v", pacientID)
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
		log.Printf("[CONSULTATIE] Error fetching consultatii by DoctorID: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var consultatii []models.Consultatie
	for cur.Next(ctx) {
		var consultatie models.Consultatie
		if err := cur.Decode(&consultatie); err != nil {
			log.Printf("[CONSULTATIE] Error decoding consultatie: %v", err)
			return nil, err
		}
		consultatii = append(consultatii, consultatie)
	}

	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATIE] Error iterating over results: %v", err)
		return nil, err
	}

	log.Printf("consultatii: %v", consultatii)

	log.Printf("[CONSULTATIE] Consultatii by DoctorID retrieved successfully: %v", doctorID)
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
		log.Printf("[CONSULTATIE] Error fetching consultatii by Date: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var consultatii []models.Consultatie
	for cur.Next(ctx) {
		var consultatie models.Consultatie
		if err := cur.Decode(&consultatie); err != nil {
			log.Printf("[CONSULTATIE] Error decoding consultatie: %v", err)
			return nil, err
		}
		consultatii = append(consultatii, consultatie)
	}

	if err := cur.Err(); err != nil {
		log.Printf("[CONSULTATIE] Error iterating over results: %v", err)
		return nil, err
	}

	log.Printf("[CONSULTATIE] Consultatii by Date retrieved successfully: %v", date)
	return consultatii, nil
}

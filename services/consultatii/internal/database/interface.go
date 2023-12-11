package database

import (
	"context"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Database interface {
	// create
	SaveConsultation(ctx context.Context, consultatie *models.Consultation) (primitive.ObjectID, error)

	// retrieve
	FetchConsultations(ctx context.Context, filter bson.M, page int, limit int) ([]models.Consultation, error)
	FetchConsultationByID(ctx context.Context, consultationID primitive.ObjectID) (*models.Consultation, error)

	// update
	UpdateConsultationByID(ctx context.Context, consultatie *models.Consultation) (int, error)

	// delete
	DeleteConsultationByID(ctx context.Context, consultationID primitive.ObjectID) (int, error)
	DeleteConsultationsByPatientOrDoctorID(ctx context.Context, id int) (int, error)

	// Add more methods as needed

	Close(ctx context.Context) error
}

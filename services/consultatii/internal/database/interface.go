package database

import (
	"context"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Database interface {
	// create
	SaveConsultation(ctx context.Context, consultatie *models.Consultation) (primitive.ObjectID, error)

	// retrieve
	FetchAllConsultations(ctx context.Context, page, limit int) ([]models.Consultation, error)
	FetchConsultationByID(ctx context.Context, consultationID primitive.ObjectID) (*models.Consultation, error)
	FetchConsultationsByPatientID(ctx context.Context, patientID int, page, limit int) ([]models.Consultation, error)
	FetchConsultationsByDoctorID(ctx context.Context, doctorID int, page, limit int) ([]models.Consultation, error)
	FetchConsultationsByDate(ctx context.Context, date time.Time, page, limit int) ([]models.Consultation, error)

	FetchConsultationsByFilter(ctx context.Context, filter bson.D, page int, limit int) ([]models.Consultation, error)

	// update
	UpdateConsultationByID(ctx context.Context, consultatie *models.Consultation) (int, error)

	// delete
	DeleteConsultationByID(ctx context.Context, consultationID primitive.ObjectID) (int, error)

	// Add more methods as needed

	Close(ctx context.Context) error
}

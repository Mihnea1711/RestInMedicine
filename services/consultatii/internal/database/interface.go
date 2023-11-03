package database

import (
	"context"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Database interface {
	// create
	SaveConsultatie(ctx context.Context, consultatie *models.Consultatie) error

	// retrieve
	FetchAllConsultatii(ctx context.Context, page, limit int) ([]models.Consultatie, error)
	FetchConsultatieByID(ctx context.Context, id primitive.ObjectID) (*models.Consultatie, error)
	FetchConsultatiiByPacientID(ctx context.Context, pacientID int, page, limit int) ([]models.Consultatie, error)
	FetchConsultatiiByDoctorID(ctx context.Context, doctorID int, page, limit int) ([]models.Consultatie, error)
	FetchConsultatiiByDate(ctx context.Context, date time.Time, page, limit int) ([]models.Consultatie, error)

	// update
	UpdateConsultatieByID(ctx context.Context, consultatie *models.Consultatie) (int64, error)

	// delete
	DeleteConsultatieByID(ctx context.Context, id primitive.ObjectID) (int64, error)

	// Add more methods as needed

	Close(ctx context.Context) error
}

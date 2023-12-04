package database

import (
	"context"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
)

type Database interface {
	SavePatient(ctx context.Context, patient *models.Pacient) (int, error)

	FetchPatients(ctx context.Context, page, limit int) ([]models.Pacient, error)
	FetchActivePatients(ctx context.Context, page, limit int) ([]models.Pacient, error)
	FetchPatientByID(ctx context.Context, patientID int) (*models.Pacient, error)
	FetchPatientByEmail(ctx context.Context, email string) (*models.Pacient, error)
	FetchPatientByUserID(ctx context.Context, userID int) (*models.Pacient, error)

	UpdatePatientByID(ctx context.Context, patient *models.Pacient) (int, error)
	DeletePatientByID(ctx context.Context, patientID int) (int, error)
	DeletePatientByUserID(ctx context.Context, patientUserID int) (int, error)

	SetPatientActivityByUserID(ctx context.Context, isActive bool, userID int) (int, error)

	Close() error
}

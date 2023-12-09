package database

import (
	"context"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
)

type Database interface {
	SavePatient(ctx context.Context, patient *models.Patient) (int, error)

	FetchPatients(ctx context.Context, filters map[string]interface{}, page, limit int) ([]models.Patient, error)
	FetchPatientByID(ctx context.Context, patientID int) (*models.Patient, error)
	FetchPatientByEmail(ctx context.Context, email string) (*models.Patient, error)
	FetchPatientByUserID(ctx context.Context, userID int) (*models.Patient, error)

	UpdatePatientByID(ctx context.Context, patient *models.Patient) (int, error)
	DeletePatientByID(ctx context.Context, patientID int) (int, error)
	DeletePatientByUserID(ctx context.Context, patientUserID int) (int, error)

	SetPatientActivityByUserID(ctx context.Context, isActive bool, userID int) (int, error)

	Close() error
}

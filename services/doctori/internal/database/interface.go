package database

import (
	"context"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
)

type Database interface {
	SaveDoctor(ctx context.Context, doctor *models.Doctor) error

	FetchDoctors(ctx context.Context) ([]models.Doctor, error)
	FetchDoctorByID(ctx context.Context, id int) (*models.Doctor, error)
	FetchDoctorByEmail(ctx context.Context, email string) (*models.Doctor, error)
	FetchDoctorByUserID(ctx context.Context, userID int) (*models.Doctor, error)

	UpdateDoctorByID(ctx context.Context, doctor *models.Doctor) error
	DeleteDoctorByID(ctx context.Context, id int) error

	// ... add more methods
}

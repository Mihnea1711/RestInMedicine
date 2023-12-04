package database

import (
	"context"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
)

type Database interface {
	SaveDoctor(ctx context.Context, doctor *models.Doctor) (int, error)

	FetchDoctors(ctx context.Context, page, limit int) ([]models.Doctor, error)
	FetchActiveDoctors(ctx context.Context, page, limit int) ([]models.Doctor, error)
	FetchDoctorByID(ctx context.Context, doctorID int) (*models.Doctor, error)
	FetchDoctorByEmail(ctx context.Context, email string) (*models.Doctor, error)
	FetchDoctorByUserID(ctx context.Context, userID int) (*models.Doctor, error)

	UpdateDoctorByID(ctx context.Context, doctor *models.Doctor) (int, error)
	DeleteDoctorByID(ctx context.Context, doctorID int) (int, error)
	DeleteDoctorByUserID(ctx context.Context, doctorUserID int) (int, error)

	SetDoctorActivityByUserID(ctx context.Context, isActive bool, userID int) (int, error)

	Close() error
}

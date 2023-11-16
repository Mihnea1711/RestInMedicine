package database

import (
	"context"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
)

type Database interface {
	SavePacient(ctx context.Context, doctor *models.Pacient) error

	FetchPacienti(ctx context.Context, page, limit int) ([]models.Pacient, error)
	FetchPacientByID(ctx context.Context, id int) (*models.Pacient, error)
	FetchPacientByEmail(ctx context.Context, email string, page, limit int) (*models.Pacient, error)
	FetchPacientByUserID(ctx context.Context, userID int) (*models.Pacient, error)

	UpdatePacientByID(ctx context.Context, doctor *models.Pacient) (int, error)
	DeletePacientByID(ctx context.Context, id int) (int, error)

	// ... add more methods

	Close() error
}

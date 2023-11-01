package database

import (
	"context"
	"time"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
)

type Database interface {
	SaveProgramare(ctx context.Context, programare *models.Programare) error

	FetchProgramari(ctx context.Context, page, limit int) ([]models.Programare, error)
	FetchProgramareByID(ctx context.Context, id int) (*models.Programare, error)
	FetchProgramariByPacientID(ctx context.Context, id, page, limit int) ([]models.Programare, error)
	FetchProgramariByDoctorID(ctx context.Context, id, page, limit int) ([]models.Programare, error)
	FetchProgramariByDate(ctx context.Context, date time.Time, page, limit int) ([]models.Programare, error)
	FetchProgramariByStatus(ctx context.Context, state string, page, limit int) ([]models.Programare, error)

	UpdateProgramareByID(ctx context.Context, programare *models.Programare) (int64, error)
	DeleteProgramareByID(ctx context.Context, id int) (int64, error)

	// add more

	Close() error
}

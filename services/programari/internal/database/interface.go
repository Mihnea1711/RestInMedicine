package database

import (
	"context"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
)

type Database interface {
	SaveAppointment(ctx context.Context, programare *models.Appointment) (int, error)

	FetchAppointmentByID(ctx context.Context, appointmentID int) (*models.Appointment, error)

	FetchAppointments(ctx context.Context, filters map[string]interface{}, page, limit int) ([]models.Appointment, error)

	UpdateAppointmentByID(ctx context.Context, programare *models.Appointment) (int, error)
	DeleteAppointmentByID(ctx context.Context, appointmentID int) (int, error)

	// add more

	Close() error
}

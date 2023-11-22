package database

import (
	"context"
	"time"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
)

type Database interface {
	SaveAppointment(ctx context.Context, programare *models.Appointment) (int, error)

	FetchAppointments(ctx context.Context, page, limit int) ([]models.Appointment, error)
	FetchAppointmentByID(ctx context.Context, appointmentID int) (*models.Appointment, error)
	FetchAppointmentsByPatientID(ctx context.Context, patientID, page, limit int) ([]models.Appointment, error)
	FetchAppointmentsByDoctorID(ctx context.Context, doctorID, page, limit int) ([]models.Appointment, error)
	FetchAppointmentsByDate(ctx context.Context, date time.Time, page, limit int) ([]models.Appointment, error)
	FetchAppointmentsByStatus(ctx context.Context, status string, page, limit int) ([]models.Appointment, error)

	UpdateAppointmentByID(ctx context.Context, programare *models.Appointment) (int, error)
	DeleteAppointmentByID(ctx context.Context, appointmentID int) (int, error)

	// add more

	Close() error
}

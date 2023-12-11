package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// FetchAppointments queries the database for appointments based on the provided filters.
// If filters is empty, it retrieves all appointments.
func (db *MySQLDatabase) FetchAppointments(ctx context.Context, filters map[string]interface{}, page, limit int) ([]models.Appointment, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	qb := squirrel.Select("*").From(utils.AppointmentTableName)

	// Check if filters is not empty and add WHERE clause if needed
	if len(filters) > 0 {
		qb = qb.Where(filters)
	}

	// Add LIMIT and OFFSET for pagination
	qb = qb.Limit(uint64(limit)).Offset(uint64(offset))

	query, args, err := qb.ToSql()
	if err != nil {
		log.Printf("[APPOINTMENT] FetchAppointments: Failed to construct SQL query: %v", err)
		return nil, fmt.Errorf("internal server error")
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("[APPOINTMENT] FetchAppointments: Failed to query database: %v", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err := rows.Scan(
			&appointment.IDProgramare,
			&appointment.IDPatient,
			&appointment.IDDoctor,
			&appointment.Date,
			&appointment.Status,
		)
		if err != nil {
			log.Printf("[APPOINTMENT] FetchAppointments: Failed to scan rows: %v", err)
			return nil, fmt.Errorf("internal server error")
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

// FetchAppointmentByID retrieves a appointment by its ID.
func (db *MySQLDatabase) FetchAppointmentByID(ctx context.Context, appointmentID int) (*models.Appointment, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", utils.AppointmentTableName, utils.ColumnIDProgramare)

	log.Printf("[APPOINTMENT] Attempting to fetch appointment with ID %d", appointmentID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, appointmentID)

	appointment := models.Appointment{}
	if err := row.Scan(
		&appointment.IDProgramare,
		&appointment.IDPatient,
		&appointment.IDDoctor,
		&appointment.Date,
		&appointment.Status,
	); err != nil {
		log.Printf("[APPOINTMENT] Error fetching appointment by ID: %v", err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched appointment by ID %d.", appointmentID)
	return &appointment, nil
}

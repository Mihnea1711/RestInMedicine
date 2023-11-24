package mysql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// FetchAppointments lists all appointments with pagination.
func (db *MySQLDatabase) FetchAppointments(ctx context.Context, page, limit int) ([]models.Appointment, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT * FROM %s LIMIT ? OFFSET ?", utils.AppointmentTableName)

	log.Printf("[APPOINTMENT] Attempting to fetch appoitments with limit=%d, offset=%d", limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch appointments: %v", err)
		return nil, err
	}
	defer rows.Close()

	appointments := []models.Appointment{}
	for rows.Next() {
		appointment := models.Appointment{}
		if err := rows.Scan(
			&appointment.IDProgramare,
			&appointment.IDPacient,
			&appointment.IDDoctor,
			&appointment.Date,
			&appointment.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning appointment: %v", err)
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d appointments.", len(appointments))
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
		&appointment.IDPacient,
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

// FetchAppointmentsByPacientID lists appointments for a specific patient with pagination.
func (db *MySQLDatabase) FetchAppointmentsByPatientID(ctx context.Context, patientID, page, limit int) ([]models.Appointment, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT ? OFFSET ?", utils.AppointmentTableName, utils.ColumnIDPacient)

	log.Printf("[APPOINTMENT] Attempting to fetch appoitments by patientID %d with limit=%d, offset=%d", patientID, limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, patientID, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch appointments by patientID: %v", err)
		return nil, err
	}
	defer rows.Close()

	appointments := []models.Appointment{}
	for rows.Next() {
		appointment := models.Appointment{}
		if err := rows.Scan(
			&appointment.IDProgramare,
			&appointment.IDPacient,
			&appointment.IDDoctor,
			&appointment.Date,
			&appointment.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning appointment by patientID: %v", err)
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error after iterating over rows by patientID: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d appointments by patientID %d.", len(appointments), patientID)
	return appointments, nil
}

// FetchAppointmentsByDoctorID lists appointments for a specific doctor with pagination.
func (db *MySQLDatabase) FetchAppointmentsByDoctorID(ctx context.Context, doctorID, page, limit int) ([]models.Appointment, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT ? OFFSET ?", utils.AppointmentTableName, utils.ColumnIDDoctor)

	log.Printf("[APPOINTMENT] Attempting to fetch appoitments by doctorID %d with limit=%d, offset=%d", doctorID, limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, doctorID, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch appointments by doctorID: %v", err)
		return nil, err
	}
	defer rows.Close()

	appointments := []models.Appointment{}
	for rows.Next() {
		appointment := models.Appointment{}
		if err := rows.Scan(
			&appointment.IDProgramare,
			&appointment.IDPacient,
			&appointment.IDDoctor,
			&appointment.Date,
			&appointment.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning appointments by doctorID: %v", err)
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error after iterating over rows by doctorID: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d appointments by doctorID %d.", len(appointments), doctorID)
	return appointments, nil
}

// FetchAppointmentsByDate lists appointments based on a specific date with pagination.
func (db *MySQLDatabase) FetchAppointmentsByDate(ctx context.Context, date time.Time, page, limit int) ([]models.Appointment, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT ? OFFSET ?", utils.AppointmentTableName, utils.ColumnDate)

	log.Printf("[APPOINTMENT] Attempting to fetch appointments by date %s with limit=%d, offset=%d", date, limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, date, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch appointments by date: %v", err)
		return nil, err
	}
	defer rows.Close()

	appointments := []models.Appointment{}
	for rows.Next() {
		appointment := models.Appointment{}
		if err := rows.Scan(
			&appointment.IDProgramare,
			&appointment.IDPacient,
			&appointment.IDDoctor,
			&appointment.Date,
			&appointment.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning appointments by date: %v", err)
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error after iterating over rows by date: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d appointments by date %s.", len(appointments), date)
	return appointments, nil
}

// FetchAppointmentsByState lists appointments based on a specific status with pagination.
func (db *MySQLDatabase) FetchAppointmentsByStatus(ctx context.Context, status string, page, limit int) ([]models.Appointment, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT ? OFFSET ?", utils.AppointmentTableName, utils.ColumnStatus)

	log.Printf("[APPOINTMENT] Attempting to fetch appointments by status %s with limit=%d, offset=%d", status, limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, status, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch appointments by status: %v", err)
		return nil, err
	}
	defer rows.Close()

	appointments := []models.Appointment{}
	for rows.Next() {
		appointment := models.Appointment{}
		if err := rows.Scan(
			&appointment.IDProgramare,
			&appointment.IDPacient,
			&appointment.IDDoctor,
			&appointment.Date,
			&appointment.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning appointments by state: %v", err)
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error scanning appointments by state: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d appointments by status %s.", len(appointments), status)
	return appointments, nil
}

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// FetchProgramari lists all programari with pagination.
func (db *MySQLDatabase) FetchProgramari(ctx context.Context, page, limit int) ([]models.Programare, error) {
	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT * FROM %s LIMIT ? OFFSET ?", utils.AppointmentTableName)

	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch programari: %v", err)
		return nil, err
	}
	defer rows.Close()

	programari := []models.Programare{}
	for rows.Next() {
		programare := models.Programare{}
		if err := rows.Scan(
			&programare.IDProgramare,
			&programare.IDPacient,
			&programare.IDDoctor,
			&programare.Date,
			&programare.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning programare: %v", err)
			return nil, err
		}
		programari = append(programari, programare)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error scanning programari: %v", err)
		return nil, err
	}

	return programari, nil
}

// FetchProgramareByID retrieves a programare by its ID.
func (db *MySQLDatabase) FetchProgramareByID(ctx context.Context, appointmentId int) (*models.Programare, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_programare = ?", utils.AppointmentTableName)

	row := db.QueryRowContext(ctx, query, appointmentId)

	programare := models.Programare{}
	if err := row.Scan(
		&programare.IDProgramare,
		&programare.IDPacient,
		&programare.IDDoctor,
		&programare.Date,
		&programare.Status,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("[APPOINTMENT] Error retrieving programare by ID: %v", err)
		return nil, err
	}

	return &programare, nil
}

// FetchProgramariByPacientID lists programari for a specific pacient with pagination.
func (db *MySQLDatabase) FetchProgramariByPacientID(ctx context.Context, patientId, page, limit int) ([]models.Programare, error) {
	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", utils.AppointmentTableName, utils.ColumnIDProgramare)

	rows, err := db.QueryContext(ctx, query, patientId, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch programari by pacient ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	programari := []models.Programare{}
	for rows.Next() {
		programare := models.Programare{}
		if err := rows.Scan(
			&programare.IDProgramare,
			&programare.IDPacient,
			&programare.IDDoctor,
			&programare.Date,
			&programare.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning programare by pacient ID: %v", err)
			return nil, err
		}
		programari = append(programari, programare)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error scanning programari by pacient ID: %v", err)
		return nil, err
	}

	return programari, nil
}

// FetchProgramariByDoctorID lists programari for a specific doctor with pagination.
func (db *MySQLDatabase) FetchProgramariByDoctorID(ctx context.Context, doctorId, page, limit int) ([]models.Programare, error) {
	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT ? OFFSET ?", utils.AppointmentTableName, utils.ColumnIDDoctor)

	rows, err := db.QueryContext(ctx, query, doctorId, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch programari by doctor ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	programari := []models.Programare{}
	for rows.Next() {
		programare := models.Programare{}
		if err := rows.Scan(
			&programare.IDProgramare,
			&programare.IDPacient,
			&programare.IDDoctor,
			&programare.Date,
			&programare.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning programari by doctor ID: %v", err)
			return nil, err
		}
		programari = append(programari, programare)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error scanning programari by doctor ID: %v", err)
		return nil, err
	}

	return programari, nil
}

// FetchProgramariByDate lists programari based on a specific date with pagination.
func (db *MySQLDatabase) FetchProgramariByDate(ctx context.Context, date time.Time, page, limit int) ([]models.Programare, error) {
	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT ? OFFSET ?", utils.AppointmentTableName, utils.ColumnDate)
	log.Printf("%s / %s / %d / %d", query, date, page, limit)

	rows, err := db.QueryContext(ctx, query, date, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch programari by date: %v", err)
		return nil, err
	}
	defer rows.Close()

	programari := []models.Programare{}
	for rows.Next() {
		programare := models.Programare{}
		if err := rows.Scan(
			&programare.IDProgramare,
			&programare.IDPacient,
			&programare.IDDoctor,
			&programare.Date,
			&programare.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning programari by date: %v", err)
			return nil, err
		}
		programari = append(programari, programare)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error scanning programari by date: %v", err)
		return nil, err
	}

	return programari, nil
}

// FetchProgramariByState lists programari based on a specific status with pagination.
func (db *MySQLDatabase) FetchProgramariByStatus(ctx context.Context, state string, page, limit int) ([]models.Programare, error) {
	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT ? OFFSET ?", utils.AppointmentTableName, utils.ColumnStatus)

	rows, err := db.QueryContext(ctx, query, state, limit, offset)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to fetch programari by state: %v", err)
		return nil, err
	}
	defer rows.Close()

	programari := []models.Programare{}
	for rows.Next() {
		programare := models.Programare{}
		if err := rows.Scan(
			&programare.IDProgramare,
			&programare.IDPacient,
			&programare.IDDoctor,
			&programare.Date,
			&programare.Status,
		); err != nil {
			log.Printf("[APPOINTMENT] Error scanning programari by state: %v", err)
			return nil, err
		}
		programari = append(programari, programare)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[APPOINTMENT] Error scanning programari by state: %v", err)
		return nil, err
	}

	return programari, nil
}

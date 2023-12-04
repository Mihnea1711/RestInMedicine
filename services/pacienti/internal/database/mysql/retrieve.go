package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) FetchPatients(ctx context.Context, page, limit int) ([]models.Pacient, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s LIMIT ? OFFSET ?",
		utils.ColumnIDPacient,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnCNP,
		utils.ColumnDataNasterii,
		utils.ColumnIsActive,
		utils.TableName,
	)

	log.Printf("[PATIENT] Attempting to fetch patients with limit=%d, offset=%d", limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to fetch patients: %v", err)
		return nil, err
	}
	defer rows.Close()

	var patients []models.Pacient
	for rows.Next() {
		var patient models.Pacient
		err := rows.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
		if err != nil {
			log.Printf("[PATIENT] Error scanning patient row: %v", err)
			return nil, err
		}
		patients = append(patients, patient)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[PATIENT] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d patients.", len(patients))
	return patients, nil
}

func (db *MySQLDatabase) FetchActivePatients(ctx context.Context, page, limit int) ([]models.Pacient, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = true LIMIT ? OFFSET ?",
		utils.ColumnIDPacient,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnCNP,
		utils.ColumnDataNasterii,
		utils.ColumnIsActive,
		utils.TableName,
		utils.ColumnIsActive,
	)

	log.Printf("[PATIENT] Attempting to fetch active patients with limit=%d, offset=%d", limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to fetch active patients: %v", err)
		return nil, err
	}
	defer rows.Close()

	var patients []models.Pacient
	for rows.Next() {
		var patient models.Pacient
		err := rows.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
		if err != nil {
			log.Printf("[PATIENT] Error scanning patient row: %v", err)
			return nil, err
		}
		patients = append(patients, patient)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[PATIENT] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d active patients.", len(patients))
	return patients, nil
}

func (db *MySQLDatabase) FetchPatientByID(ctx context.Context, patientID int) (*models.Pacient, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDPacient,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnCNP,
		utils.ColumnDataNasterii,
		utils.ColumnIsActive,
		utils.TableName,
		utils.ColumnIDPacient,
	)

	log.Printf("[PATIENT] Attempting to fetch patient with ID %d", patientID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, patientID)

	var patient models.Pacient
	err := row.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
	if err != nil {
		log.Printf("[PATIENT] Error fetching patient by ID %d: %v", patientID, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by ID %d.", patientID)
	return &patient, nil
}

func (db *MySQLDatabase) FetchPatientByEmail(ctx context.Context, email string) (*models.Pacient, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDPacient,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnCNP,
		utils.ColumnDataNasterii,
		utils.ColumnIsActive,
		utils.TableName,
		utils.ColumnEmail,
	)
	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, email)

	log.Printf("[PATIENT] Attempting to fetch patient by email %s", email)

	var patient models.Pacient
	err := row.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
	if err != nil {
		log.Printf("[PATIENT] Error fetching patient by email %s: %v", email, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by email %s.", email)
	return &patient, nil
}

func (db *MySQLDatabase) FetchPatientByUserID(ctx context.Context, userID int) (*models.Pacient, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDPacient,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnCNP,
		utils.ColumnDataNasterii,
		utils.ColumnIsActive,
		utils.TableName,
		utils.ColumnIDUser,
	)

	log.Printf("[PATIENT] Attempting to fetch patient by user ID %d", userID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, userID)

	var patient models.Pacient
	err := row.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
	if err != nil {
		log.Printf("[PATIENT] Error fetching patient by user ID %d: %v", userID, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by user ID %d.", userID)
	return &patient, nil
}

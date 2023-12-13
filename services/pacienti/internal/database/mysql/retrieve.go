package mysql

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) FetchPatients(ctx context.Context, filters map[string]interface{}, page, limit int) ([]models.Patient, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	qb := squirrel.Select("*").From(utils.PatientTableName)

	// Check if filters is not empty and add WHERE clause if needed
	if len(filters) > 0 {
		// Check if filters contain 'first_name'
		if firstName, ok := filters[utils.ColumnFirstName].(string); ok && firstName != "" {
			// Apply the LIKE filter for the 'first_name' column
			qb = qb.Where(squirrel.Like{utils.ColumnFirstName: "%" + strings.ToLower(firstName) + "%"})
		}

		// Check if filters contain 'isActive'
		if isActive, ok := filters[utils.ColumnIsActive].(bool); ok {
			// Apply the filter for 'isActive'
			qb = qb.Where(squirrel.Eq{utils.ColumnIsActive: isActive})
		}
	}

	log.Printf("[PATIENT] Attempting to fetch patients with limit=%d, offset=%d", limit, offset)

	// Add LIMIT and OFFSET for pagination
	qb = qb.Limit(uint64(limit)).Offset(uint64(offset))

	query, args, err := qb.ToSql()
	if err != nil {
		log.Printf("[DOCTOR] FetchPatients: Failed to construct SQL query: %v", err)
		return nil, fmt.Errorf("internal server error")
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("[DOCTOR] FetchPatients: Failed to query database: %v", err)
		return nil, fmt.Errorf("failed to fetch patients: %v", err)
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		err := rows.Scan(&patient.IDPatient, &patient.IDUser, &patient.FirstName, &patient.SecondName, &patient.Email, &patient.PhoneNumber, &patient.CNP, &patient.BirthDay, &patient.IsActive)
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

func (db *MySQLDatabase) FetchPatientByID(ctx context.Context, patientID int) (*models.Patient, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDPatient,
		utils.ColumnIDUser,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnCNP,
		utils.ColumnBirthDay,
		utils.ColumnIsActive,
		utils.PatientTableName,
		utils.ColumnIDPatient,
	)

	log.Printf("[PATIENT] Attempting to fetch patient with ID %d", patientID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, patientID)

	var patient models.Patient
	err := row.Scan(&patient.IDPatient, &patient.IDUser, &patient.FirstName, &patient.SecondName, &patient.Email, &patient.PhoneNumber, &patient.CNP, &patient.BirthDay, &patient.IsActive)
	if err != nil {
		log.Printf("[PATIENT] Error fetching patient by ID %d: %v", patientID, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by ID %d.", patientID)
	return &patient, nil
}

func (db *MySQLDatabase) FetchPatientByEmail(ctx context.Context, email string) (*models.Patient, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDPatient,
		utils.ColumnIDUser,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnCNP,
		utils.ColumnBirthDay,
		utils.ColumnIsActive,
		utils.PatientTableName,
		utils.ColumnEmail,
	)
	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, email)

	log.Printf("[PATIENT] Attempting to fetch patient by email %s", email)

	var patient models.Patient
	err := row.Scan(&patient.IDPatient, &patient.IDUser, &patient.FirstName, &patient.SecondName, &patient.Email, &patient.PhoneNumber, &patient.CNP, &patient.BirthDay, &patient.IsActive)
	if err != nil {
		log.Printf("[PATIENT] Error fetching patient by email %s: %v", email, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by email %s.", email)
	return &patient, nil
}

func (db *MySQLDatabase) FetchPatientByUserID(ctx context.Context, userID int) (*models.Patient, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDPatient,
		utils.ColumnIDUser,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnCNP,
		utils.ColumnBirthDay,
		utils.ColumnIsActive,
		utils.PatientTableName,
		utils.ColumnIDUser,
	)

	log.Printf("[PATIENT] Attempting to fetch patient by user ID %d", userID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, userID)

	var patient models.Patient
	err := row.Scan(&patient.IDPatient, &patient.IDUser, &patient.FirstName, &patient.SecondName, &patient.Email, &patient.PhoneNumber, &patient.CNP, &patient.BirthDay, &patient.IsActive)
	if err != nil {
		log.Printf("[PATIENT] Error fetching patient by user ID %d: %v", userID, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by user ID %d.", userID)
	return &patient, nil
}

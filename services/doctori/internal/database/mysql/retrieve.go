package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) FetchDoctors(ctx context.Context, page, limit int) ([]models.Doctor, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s LIMIT ? OFFSET ?",
		utils.ColumnIDDoctor,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnSpecializare,
		utils.DoctorTableName,
	)

	log.Printf("[DOCTOR] Attempting to fetch doctors with limit=%d, offset=%d", limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Printf("[DOCTOR] Error executing query to fetch doctors: %v", err)
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor
	for rows.Next() {
		var doctor models.Doctor
		err := rows.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)
		if err != nil {
			log.Printf("[DOCTOR] Error scanning doctor row: %v", err)
			return nil, err
		}
		doctors = append(doctors, doctor)
	}

	// Check for errors from iterating over rows.
	err = rows.Err()
	if err != nil {
		log.Printf("[DOCTOR] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched %d doctors.", len(doctors))
	return doctors, nil
}

func (db *MySQLDatabase) FetchActiveDoctors(ctx context.Context, page, limit int) ([]models.Doctor, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = true LIMIT ? OFFSET ?",
		utils.ColumnIDDoctor,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnSpecializare,
		utils.ColumnIsActive,
		utils.DoctorTableName,
		utils.ColumnIsActive,
	)

	log.Printf("[DOCTOR] Attempting to fetch active doctors with limit=%d, offset=%d", limit, offset)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Printf("[DOCTOR] Error executing query to fetch active doctors: %v", err)
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor
	for rows.Next() {
		var doctor models.Doctor
		err := rows.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare, &doctor.IsActive)
		if err != nil {
			log.Printf("[DOCTOR] Error scanning doctor row: %v", err)
			return nil, err
		}
		doctors = append(doctors, doctor)
	}

	// Check for errors from iterating over rows.
	err = rows.Err()
	if err != nil {
		log.Printf("[DOCTOR] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched %d active doctors.", len(doctors))
	return doctors, nil
}

func (db *MySQLDatabase) FetchDoctorByID(ctx context.Context, doctorID int) (*models.Doctor, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDDoctor,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnSpecializare,
		utils.ColumnIsActive,
		utils.DoctorTableName,
		utils.ColumnIDDoctor,
	)

	log.Printf("[PATIENT] Attempting to fetch doctor with ID %d", doctorID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, doctorID)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare, &doctor.IsActive)
	if err != nil {
		log.Printf("[DOCTOR] Error fetching doctor by ID %d: %v", doctorID, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by ID %d.", doctorID)
	return &doctor, nil
}

func (db *MySQLDatabase) FetchDoctorByEmail(ctx context.Context, email string) (*models.Doctor, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDDoctor,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnSpecializare,
		utils.ColumnIsActive,
		utils.DoctorTableName,
		utils.ColumnEmail,
	)

	log.Printf("[PATIENT] Attempting to fetch doctor with email %s", email)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, email)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare, &doctor.IsActive)
	if err != nil {
		log.Printf("[DOCTOR] Error fetching doctor by email %s: %v", email, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by email %s.", email)
	return &doctor, nil
}

func (db *MySQLDatabase) FetchDoctorByUserID(ctx context.Context, userID int) (*models.Doctor, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDDoctor,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnSpecializare,
		utils.ColumnIsActive,
		utils.DoctorTableName,
		utils.ColumnIDUser,
	)

	log.Printf("[PATIENT] Attempting to fetch doctor with user ID %d", userID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, userID)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare, &doctor.IsActive)
	if err != nil {
		log.Printf("[DOCTOR] Error fetching doctor by user ID %d: %v", userID, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by user ID %d.", userID)
	return &doctor, nil
}

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) FetchDoctors(ctx context.Context) ([]models.Doctor, error) {
	query := fmt.Sprintf(`SELECT id_doctor, id_user, nume, prenume, email, telefon, specializare FROM %s`, utils.DOCTOR_TABLE)

	// Execute the SQL query with context
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error executing query to fetch doctors: %v", err)
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor

	for rows.Next() {
		var doctor models.Doctor
		err := rows.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)
		if err != nil {
			log.Printf("Error scanning doctor row: %v", err)
			return nil, err
		}
		doctors = append(doctors, doctor)
	}

	// Check for errors from iterating over rows.
	err = rows.Err()
	if err != nil {
		log.Printf("Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("Successfully fetched %d doctors.", len(doctors))
	return doctors, nil
}

func (db *MySQLDatabase) FetchDoctorByID(ctx context.Context, id int) (*models.Doctor, error) {
	query := fmt.Sprintf(`SELECT id_doctor, id_user, nume, prenume, email, telefon, specializare FROM %s WHERE id_doctor = ?`, utils.DOCTOR_TABLE)
	row := db.QueryRowContext(ctx, query, id)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)
	if err != nil {
		if err == sql.ErrNoRows {
			// Not found
			return nil, nil
		}
		return nil, err
	}

	return &doctor, nil
}

func (db *MySQLDatabase) FetchDoctorByEmail(ctx context.Context, email string) (*models.Doctor, error) {
	// Define the SQL query to retrieve a doctor based on the email
	query := fmt.Sprintf(`SELECT id_doctor, id_user, nume, prenume, email, telefon, specializare FROM %s WHERE email = ?`, utils.DOCTOR_TABLE)

	// Execute the SQL query with the provided email
	row := db.QueryRowContext(ctx, query, email)

	// Create an instance of the Doctor model to hold the retrieved data
	var doctor models.Doctor

	// Scan the result into the Doctor instance
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)

	if err != nil {
		if err == sql.ErrNoRows {
			// If no results are returned
			return nil, nil
		}
		// Return any other errors
		return nil, err
	}

	// Return the retrieved doctor
	return &doctor, nil
}

func (db *MySQLDatabase) FetchDoctorByUserID(ctx context.Context, userID int) (*models.Doctor, error) {
	// Define the SQL query to retrieve a doctor based on the id_user
	query := fmt.Sprintf(`SELECT id_doctor, id_user, nume, prenume, email, telefon, specializare FROM %s WHERE id_user = ?`, utils.DOCTOR_TABLE)

	// Execute the SQL query with the provided userID
	row := db.QueryRowContext(ctx, query, userID)

	// Create an instance of the Doctor model to hold the retrieved data
	var doctor models.Doctor

	// Scan the result into the Doctor instance
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)

	if err != nil {
		if err == sql.ErrNoRows {
			// If no results are returned
			return nil, nil
		}
		// Return any other errors
		return nil, err
	}

	// Return the retrieved doctor
	return &doctor, nil
}

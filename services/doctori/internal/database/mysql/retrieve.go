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

func (db *MySQLDatabase) FetchDoctorByID(ctx context.Context, id int) (*models.Doctor, error) {
	query := fmt.Sprintf(`SELECT id_doctor, id_user, nume, prenume, email, telefon, specializare FROM %s WHERE id_doctor = ?`, utils.DOCTOR_TABLE)
	row := db.QueryRowContext(ctx, query, id)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)
	if err != nil {
		if err == sql.ErrNoRows {
			// Not found
			log.Printf("[DOCTOR] Doctor with ID %d not found.", id)
			return nil, nil
		}
		log.Printf("[DOCTOR] Error fetching doctor by ID %d: %v", id, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by ID %d.", id)
	return &doctor, nil
}

func (db *MySQLDatabase) FetchDoctorByEmail(ctx context.Context, email string) (*models.Doctor, error) {
	query := fmt.Sprintf(`SELECT id_doctor, id_user, nume, prenume, email, telefon, specializare FROM %s WHERE email = ?`, utils.DOCTOR_TABLE)
	row := db.QueryRowContext(ctx, query, email)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)
	if err != nil {
		if err == sql.ErrNoRows {
			// Not found
			log.Printf("[DOCTOR] Doctor with email %s not found.", email)
			return nil, nil
		}
		log.Printf("[DOCTOR] Error fetching doctor by email %s: %v", email, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by email %s.", email)
	return &doctor, nil
}

func (db *MySQLDatabase) FetchDoctorByUserID(ctx context.Context, userID int) (*models.Doctor, error) {
	query := fmt.Sprintf(`SELECT id_doctor, id_user, nume, prenume, email, telefon, specializare FROM %s WHERE id_user = ?`, utils.DOCTOR_TABLE)
	row := db.QueryRowContext(ctx, query, userID)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)
	if err != nil {
		if err == sql.ErrNoRows {
			// Not found
			log.Printf("[DOCTOR] Doctor with user ID %d not found.", userID)
			return nil, nil
		}
		log.Printf("[DOCTOR] Error fetching doctor by user ID %d: %v", userID, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by user ID %d.", userID)
	return &doctor, nil
}

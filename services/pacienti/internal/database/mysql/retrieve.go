package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) FetchPatients(ctx context.Context, page, limit int) ([]models.Pacient, error) {
	offset := (page - 1) * limit

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

	log.Printf("[PATIENT] Attempting to fetch pacienti with limit=%d, offset=%d", limit, offset)

	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to fetch pacienti: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pacients []models.Pacient

	for rows.Next() {
		var patient models.Pacient
		err := rows.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
		if err != nil {
			log.Printf("[PATIENT] Error scanning patient row: %v", err)
			return nil, err
		}
		pacients = append(pacients, patient)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[PATIENT] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d pacienti.", len(pacients))
	return pacients, nil
}

func (db *MySQLDatabase) FetchPatientByID(ctx context.Context, id int) (*models.Pacient, error) {
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
	row := db.QueryRowContext(ctx, query, id)

	log.Printf("[PATIENT] Attempting to fetch patient with ID %d", id)

	var patient models.Pacient
	err := row.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[PATIENT] Pacient with ID %d not found.", id)
			return nil, nil
		}
		log.Printf("[PATIENT] Error fetching patient by ID %d: %v", id, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by ID %d.", id)
	return &patient, nil
}

func (db *MySQLDatabase) FetchPatientByEmail(ctx context.Context, email string, page, limit int) (*models.Pacient, error) {
	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ? LIMIT ? OFFSET ?",
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
	row := db.QueryRowContext(ctx, query, email, limit, offset)

	log.Printf("[PATIENT] Attempting to fetch patient by email %s with limit=%d, offset=%d", email, limit, offset)

	var patient models.Pacient
	err := row.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[PATIENT] Patient with email %s not found.", email)
			return nil, nil
		}
		log.Printf("[PATIENT] Error fetching patient by email %s: %v", email, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by email %s.", email)
	return &patient, nil
}

func (db *MySQLDatabase) FetchPatientByUserID(ctx context.Context, userID int) (*models.Pacient, error) {
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
	row := db.QueryRowContext(ctx, query, userID)

	log.Printf("[PATIENT] Attempting to fetch patient by user ID %d", userID)

	var patient models.Pacient
	err := row.Scan(&patient.IDPacient, &patient.IDUser, &patient.Nume, &patient.Prenume, &patient.Email, &patient.Telefon, &patient.CNP, &patient.DataNasterii, &patient.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[PATIENT] Patient with user ID %d not found.", userID)
			return nil, nil
		}
		log.Printf("[PATIENT] Error fetching patient by user ID %d: %v", userID, err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched patient by user ID %d.", userID)
	return &patient, nil
}

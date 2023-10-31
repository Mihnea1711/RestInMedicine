package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
)

func (db *MySQLDatabase) FetchPacienti(ctx context.Context) ([]models.Pacient, error) {
	query := `SELECT cnp, id_user, nume, prenume, email, telefon, data_nasterii, is_active FROM pacient`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[PACIENTI] Error executing query to fetch pacienti: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pacients []models.Pacient

	for rows.Next() {
		var pacient models.Pacient
		err := rows.Scan(&pacient.CNP, &pacient.IDUser, &pacient.Nume, &pacient.Prenume, &pacient.Email, &pacient.Telefon, &pacient.DataNasterii, &pacient.IsActive)
		if err != nil {
			log.Printf("[PACIENTI] Error scanning pacient row: %v", err)
			return nil, err
		}
		pacients = append(pacients, pacient)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[PACIENTI] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[PACIENTI] Successfully fetched %d pacienti.", len(pacients))
	return pacients, nil
}

func (db *MySQLDatabase) FetchPacientByID(ctx context.Context, id int) (*models.Pacient, error) {
	return nil, nil
	// TODO
}

func (db *MySQLDatabase) FetchPacientByEmail(ctx context.Context, email string) (*models.Pacient, error) {
	query := `SELECT cnp, id_user, nume, prenume, email, telefon, data_nasterii, is_active FROM pacient WHERE email = ?`
	row := db.QueryRowContext(ctx, query, email)

	var pacient models.Pacient
	err := row.Scan(&pacient.CNP, &pacient.IDUser, &pacient.Nume, &pacient.Prenume, &pacient.Email, &pacient.Telefon, &pacient.DataNasterii, &pacient.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[PACIENTI] Pacient with email %s not found.", email)
			return nil, nil
		}
		log.Printf("[PACIENTI] Error fetching pacient by email %s: %v", email, err)
		return nil, err
	}

	log.Printf("[PACIENTI] Fetched pacient by email %s successfully.", email)
	return &pacient, nil
}

func (db *MySQLDatabase) FetchPacientByUserID(ctx context.Context, userID int) (*models.Pacient, error) {
	query := `SELECT cnp, id_user, nume, prenume, email, telefon, data_nasterii, is_active FROM pacient WHERE id_user = ?`
	row := db.QueryRowContext(ctx, query, userID)

	var pacient models.Pacient
	err := row.Scan(&pacient.CNP, &pacient.IDUser, &pacient.Nume, &pacient.Prenume, &pacient.Email, &pacient.Telefon, &pacient.DataNasterii, &pacient.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[PACIENTI] Pacient with user ID %d not found.", userID)
			return nil, nil
		}
		log.Printf("[PACIENTI] Error fetching pacient by user ID %d: %v", userID, err)
		return nil, err
	}

	log.Printf("[PACIENTI] Fetched pacient by user ID %d successfully.", userID)
	return &pacient, nil
}

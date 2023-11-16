package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) FetchPacienti(ctx context.Context) ([]models.Pacient, error) {
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s",
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

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to fetch pacienti: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pacients []models.Pacient

	for rows.Next() {
		var pacient models.Pacient
		err := rows.Scan(&pacient.IDPacient, &pacient.IDUser, &pacient.Nume, &pacient.Prenume, &pacient.Email, &pacient.Telefon, &pacient.CNP, &pacient.DataNasterii, &pacient.IsActive)
		if err != nil {
			log.Printf("[PATIENT] Error scanning pacient row: %v", err)
			return nil, err
		}
		pacients = append(pacients, pacient)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[PATIENT] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[PATIENT] Successfully fetched %d pacienti.", len(pacients))
	return pacients, nil
}

func (db *MySQLDatabase) FetchPacientByID(ctx context.Context, id int) (*models.Pacient, error) {
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

	var pacient models.Pacient
	err := row.Scan(&pacient.IDPacient, &pacient.IDUser, &pacient.Nume, &pacient.Prenume, &pacient.Email, &pacient.Telefon, &pacient.CNP, &pacient.DataNasterii, &pacient.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[PATIENT] Pacient with ID %d not found.", id)
			return nil, nil
		}
		log.Printf("[PATIENT] Error fetching pacient by ID %d: %v", id, err)
		return nil, err
	}

	log.Printf("[PATIENT] Fetched pacient by ID %d successfully.", id)
	return &pacient, nil
}

func (db *MySQLDatabase) FetchPacientByEmail(ctx context.Context, email string) (*models.Pacient, error) {
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
	row := db.QueryRowContext(ctx, query, email)

	var pacient models.Pacient
	err := row.Scan(&pacient.IDPacient, &pacient.IDUser, &pacient.Nume, &pacient.Prenume, &pacient.Email, &pacient.Telefon, &pacient.CNP, &pacient.DataNasterii, &pacient.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[PATIENT] Pacient with email %s not found.", email)
			return nil, nil
		}
		log.Printf("[PATIENT] Error fetching pacient by email %s: %v", email, err)
		return nil, err
	}

	log.Printf("[PATIENT] Fetched pacient by email %s successfully.", email)
	return &pacient, nil
}

func (db *MySQLDatabase) FetchPacientByUserID(ctx context.Context, userID int) (*models.Pacient, error) {
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

	var pacient models.Pacient
	err := row.Scan(&pacient.IDPacient, &pacient.IDUser, &pacient.Nume, &pacient.Prenume, &pacient.Email, &pacient.Telefon, &pacient.CNP, &pacient.DataNasterii, &pacient.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[PATIENT] Pacient with user ID %d not found.", userID)
			return nil, nil
		}
		log.Printf("[PATIENT] Error fetching pacient by user ID %d: %v", userID, err)
		return nil, err
	}

	log.Printf("[PATIENT] Fetched pacient by user ID %d successfully.", userID)
	return &pacient, nil
}

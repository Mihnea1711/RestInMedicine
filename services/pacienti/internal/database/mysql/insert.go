package mysql

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
)

func (db *MySQLDatabase) SavePacient(ctx context.Context, pacient *models.Pacient) error {
	// Construct the SQL insert query
	query := `INSERT INTO pacient (id_user, nume, prenume, email, telefon, cnp, data_nasterii, is_active) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the SQL statement
	_, err := db.ExecContext(ctx, query, pacient.IDUser, pacient.Nume, pacient.Prenume, pacient.Email, pacient.Telefon, pacient.CNP, pacient.DataNasterii, pacient.IsActive)
	if err != nil {
		log.Printf("[PACIENT] Error executing query to save pacient: %v", err)
		return err
	}

	log.Println("[PACIENT] Pacient saved successfully.")
	return nil
}

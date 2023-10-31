package mysql

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
)

func (db *MySQLDatabase) SavePacient(ctx context.Context, pacient *models.Pacient) error {
	// Construct the SQL insert query
	query := `INSERT INTO pacient (cnp, id_user, nume, prenume, email, telefon, data_nasterii, is_active) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the SQL statement
	_, err := db.ExecContext(ctx, query, pacient.CNP, pacient.IDUser, pacient.Nume, pacient.Prenume, pacient.Email, pacient.Telefon, pacient.DataNasterii, pacient.IsActive)
	if err != nil {
		log.Printf("[PACIENTI] Error executing query to save pacient: %v", err)
		return err
	}

	log.Println("[PACIENTI] Pacient saved successfully.")
	return nil
}

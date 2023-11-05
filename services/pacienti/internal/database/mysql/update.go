package mysql

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
)

func (db *MySQLDatabase) UpdatePacientByID(ctx context.Context, pacient *models.Pacient) (int64, error) {
	// Construct the SQL update query
	query := `UPDATE pacient SET id_user=?, nume=?, prenume=?, email=?, telefon=?, cnp=?, data_nasterii=?, is_active=? WHERE id_pacient=?`

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, pacient.IDUser, pacient.Nume, pacient.Prenume, pacient.Email, pacient.Telefon, pacient.CNP, pacient.DataNasterii, pacient.IsActive, pacient.IDPacient)
	if err != nil {
		log.Printf("[PACIENT] Error executing query to update pacient with ID %d: %v", pacient.IDPacient, err)
		return 0, err
	}

	// Get the number of rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[PACIENT] Error fetching rows affected for update query of pacient with ID %d: %v", pacient.IDPacient, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[PACIENT] No pacient found with ID %d to update.", pacient.IDPacient)
	} else {
		log.Printf("[PACIENT] Successfully updated pacient with ID %d.", pacient.IDPacient)
	}

	return rowsAffected, nil
}

package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) UpdatePacientByID(ctx context.Context, pacient *models.Pacient) (int, error) {
	// Construct the SQL update query
	// query := `UPDATE pacient SET nume=?, prenume=?, email=?, telefon=?, cnp=?, data_nasterii=?, is_active=? WHERE id_pacient=?`
	query := fmt.Sprintf("UPDATE %s SET %s=?, %s=?, %s=?, %s=?, %s=?, %s=?, %s=? WHERE %s=?",
		utils.TableName,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnCNP,
		utils.ColumnDataNasterii,
		utils.ColumnIsActive,
		utils.ColumnIDPacient,
	)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, pacient.Nume, pacient.Prenume, pacient.Email, pacient.Telefon, pacient.CNP, pacient.DataNasterii, pacient.IsActive, pacient.IDPacient)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to update pacient with ID %d: %v", pacient.IDPacient, err)
		return 0, err
	}

	// Get the number of rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[PATIENT] Error fetching rows affected for update query of pacient with ID %d: %v", pacient.IDPacient, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[PATIENT] No pacient found with ID %d to update.", pacient.IDPacient)
	} else {
		log.Printf("[PATIENT] Successfully updated pacient with ID %d.", pacient.IDPacient)
	}

	return int(rowsAffected), nil
}

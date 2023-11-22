package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) UpdatePatientByID(ctx context.Context, patient *models.Pacient) (int, error) {
	// Construct the SQL update query
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

	log.Printf("[PATIENT] Attempting to update patient with ID %d: %+v", patient.IDPacient, patient)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, patient.Nume, patient.Prenume, patient.Email, patient.Telefon, patient.CNP, patient.DataNasterii, patient.IsActive, patient.IDPacient)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to update patient with ID %d: %v", patient.IDPacient, err)
		return 0, err
	}

	// Get the number of rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[PATIENT] Error fetching rows affected for update query of patient with ID %d: %v", patient.IDPacient, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[PATIENT] No pacient found with ID %d to update.", patient.IDPacient)
	} else {
		log.Printf("[PATIENT] Successfully updated pacient with ID %d. Rows affected: %d", patient.IDPacient, rowsAffected)
	}

	return int(rowsAffected), nil
}

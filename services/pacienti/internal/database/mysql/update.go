package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) UpdatePatientByID(ctx context.Context, patient *models.Patient) (int, error) {
	// Construct the SQL update query
	query := fmt.Sprintf("UPDATE %s SET %s=?, %s=?, %s=?, %s=?, %s=?, %s=?, %s=? WHERE %s=?",
		utils.PatientTableName,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnCNP,
		utils.ColumnBirthDay,
		utils.ColumnIsActive,
		utils.ColumnIDPatient,
	)

	log.Printf("[PATIENT] Attempting to update patient with ID %d", patient.IDPatient)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, patient.FirstName, patient.SecondName, patient.Email, patient.PhoneNumber, patient.CNP, patient.BirthDay, patient.IsActive, patient.IDPatient)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to update patient with ID %d: %v", patient.IDPatient, err)
		return 0, err
	}

	// Get the number of rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[PATIENT] Error fetching rows affected for update query of patient with ID %d: %v", patient.IDPatient, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[PATIENT] No patient found with ID %d to update.", patient.IDPatient)
	} else {
		log.Printf("[PATIENT] Successfully updated patient with ID %d. Rows affected: %d", patient.IDPatient, rowsAffected)
	}

	return int(rowsAffected), nil
}

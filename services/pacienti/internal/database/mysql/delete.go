package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) DeletePatientByID(ctx context.Context, patientID int) (int, error) {
	// Construct the SQL delete query
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", utils.TableName, utils.ColumnIDPacient)

	log.Printf("[PATIENT] Attempting to delete patient with ID %d", patientID)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, patientID)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to delete patient with ID %d: %v", patientID, err)
		return 0, err
	}

	// Get the number of rows affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[PATIENT] Error fetching rows affected for delete query of patient with ID %d: %v", patientID, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[PATIENT] No patient found with ID %d to delete.", patientID)
	} else {
		log.Printf("[PATIENT] Successfully deleted patient with ID %d. Rows affected: %d", patientID, rowsAffected)
	}

	return int(rowsAffected), nil
}

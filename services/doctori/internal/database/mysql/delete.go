package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) DeleteDoctorByID(ctx context.Context, doctorID int) (int, error) {
	// Construct the SQL delete query
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", utils.DoctorTableName, utils.ColumnIDDoctor)

	log.Printf("[DOCTOR] Attempting to delete doctor with doctorID %d...", doctorID)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, doctorID)
	if err != nil {
		log.Printf("[DOCTOR] Error executing query to delete doctor with doctorID %d: %v", doctorID, err)
		return 0, err
	}

	// Get the number of rows affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[DOCTOR] Error fetching rows affected for delete query of doctor with doctorID %d: %v", doctorID, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[DOCTOR] No doctor found with doctorID %d to delete.", doctorID)
	} else {
		log.Printf("[DOCTOR] Successfully deleted doctor with doctorID %d.", doctorID)
	}

	return int(rowsAffected), nil
}

func (db *MySQLDatabase) DeleteDoctorByUserID(ctx context.Context, doctorUserID int) (int, error) {
	// Construct the SQL delete query
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", utils.DoctorTableName, utils.ColumnIDUser)

	log.Printf("[DOCTOR] Attempting to delete doctor with doctorUserID %d...", doctorUserID)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, doctorUserID)
	if err != nil {
		log.Printf("[DOCTOR] Error executing query to delete doctor with doctorUserID %d: %v", doctorUserID, err)
		return 0, err
	}

	// Get the number of rows affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[DOCTOR] Error fetching rows affected for delete query of doctor with doctorUserID %d: %v", doctorUserID, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[DOCTOR] No doctor found with doctorUserID %d to delete.", doctorUserID)
	} else {
		log.Printf("[DOCTOR] Successfully deleted doctor with doctorUserID %d.", doctorUserID)
	}

	return int(rowsAffected), nil
}

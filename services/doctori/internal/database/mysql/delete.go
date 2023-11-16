package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) DeleteDoctorByID(ctx context.Context, id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", utils.DoctorTableName, utils.ColumnIDDoctor)

	log.Printf("[DOCTOR] Executing delete query for doctor with ID %d...", id)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("[DOCTOR] Error deleting doctor with ID %d: %v", id, err)
		return 0, err
	}

	// Get the number of rows affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[DOCTOR] Error fetching rows affected for delete query of doctor with ID %d: %v", id, err)
		return 0, err
	}

	return int(rowsAffected), nil
}

package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) SetDoctorActivityByUserID(ctx context.Context, isActive bool, userID int) (int, error) {
	// Construct the SQL update query
	query := fmt.Sprintf("UPDATE %s SET %s=? WHERE %s=?",
		utils.DoctorTableName,
		utils.ColumnIsActive,
		utils.ColumnIDUser,
	)

	log.Printf("[DOCTOR] Attempting to update doctor with user ID %d", userID)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, isActive, userID)
	if err != nil {
		log.Printf("[DOCTOR] Error executing query to update doctor with user ID %d: %v", userID, err)
		return 0, err
	}

	// Get the number of rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[DOCTOR] Error fetching rows affected for update query of doctor with user ID %d: %v", userID, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[DOCTOR] No doctor found with user ID %d to update.", userID)
	} else {
		log.Printf("[DOCTOR] Successfully updated doctor with user ID %d. Rows affected: %d", userID, rowsAffected)
	}

	return int(rowsAffected), nil
}

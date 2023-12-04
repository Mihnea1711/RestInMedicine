package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) SetPatientActivityByUserID(ctx context.Context, isActive bool, userID int) (int, error) {
	// Construct the SQL update query
	query := fmt.Sprintf("UPDATE %s SET %s=? WHERE %s=?",
		utils.TableName,
		utils.ColumnIsActive,
		utils.ColumnIDUser,
	)

	log.Printf("[PATIENT] Attempting to update patient with user ID %d", userID)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, isActive, userID)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to update patient with user ID %d: %v", userID, err)
		return 0, err
	}

	// Get the number of rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[PATIENT] Error fetching rows affected for update query of patient with user ID %d: %v", userID, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[PATIENT] No patient found with user ID %d to update.", userID)
	} else {
		log.Printf("[PATIENT] Successfully updated patient with user ID %d. Rows affected: %d", userID, rowsAffected)
	}

	return int(rowsAffected), nil
}

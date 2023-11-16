package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) DeletePacientByID(ctx context.Context, id int) (int, error) {
	// Construct the SQL delete query
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", utils.TableName, utils.ColumnIDPacient)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to delete pacient with ID %d: %v", id, err)
		return 0, err
	}

	// Get the number of rows affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[PATIENT] Error fetching rows affected for delete query of pacient with ID %d: %v", id, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[PATIENT] No pacient found with ID %d to delete.", id)
	} else {
		log.Printf("[PATIENT] Successfully deleted pacient with ID %d.", id)
	}

	return int(rowsAffected), nil
}

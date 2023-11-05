package mysql

import (
	"context"
	"log"
)

func (db *MySQLDatabase) DeletePacientByID(ctx context.Context, id int) (int64, error) {
	// Construct the SQL delete query
	query := `DELETE FROM pacient WHERE id_pacient = ?`

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("[PACIENT] Error executing query to delete pacient with ID %d: %v", id, err)
		return 0, err
	}

	// Get the number of rows affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[PACIENT] Error fetching rows affected for delete query of pacient with ID %d: %v", id, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[PACIENT] No pacient found with ID %d to delete.", id)
	} else {
		log.Printf("[PACIENT] Successfully deleted pacient with ID %d.", id)
	}

	return rowsAffected, nil
}

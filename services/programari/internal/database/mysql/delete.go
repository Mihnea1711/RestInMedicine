package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) DeleteAppointmentByID(ctx context.Context, appointmentID int) (int, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", utils.AppointmentTableName, utils.ColumnIDProgramare)

	log.Printf("[APPOINTMENT] Attempting to delete appointment with ID %d", appointmentID)

	// Execute the SQL statement
	res, err := db.ExecContext(ctx, query, appointmentID)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to delete appointment: %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("[APPOINTMENT] Error fetching rows affected for delete query of appointment with ID %d: %v", appointmentID, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[APPOINTMENT] No appointment found with ID %d to delete.", appointmentID)
	} else {
		log.Printf("[APPOINTMENT] Successfully deleted appointment with ID %d. Rows affected: %d", appointmentID, rowsAffected)
	}

	return int(rowsAffected), nil
}

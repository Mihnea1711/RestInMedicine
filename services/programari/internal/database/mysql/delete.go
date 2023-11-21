package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) DeleteProgramareByID(ctx context.Context, id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", utils.AppointmentTableName, utils.ColumnIDProgramare)
	res, err := db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to delete appointment: %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("[APPOINTMENT] Error getting rows affected: %v", err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[APPOINTMENT] No appointment found with ID %d to delete.", id)
	} else {
		log.Printf("[APPOINTMENT] Successfully deleted appointment with ID %d.", id)
	}

	return int(rowsAffected), nil
}

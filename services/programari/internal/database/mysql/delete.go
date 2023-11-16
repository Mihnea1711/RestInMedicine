package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) DeleteProgramareByID(ctx context.Context, id int) (int64, error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id_programare = ?`, utils.PROGRAMARE_TABLE)

	res, err := db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to delete programare: %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("[APPOINTMENT] Error getting rows affected: %v", err)
		return 0, err
	}

	log.Printf("[APPOINTMENT] %d programare deleted successfully.", rowsAffected)
	return rowsAffected, nil
}

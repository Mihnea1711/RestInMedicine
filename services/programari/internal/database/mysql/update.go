package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) UpdateProgramareByID(ctx context.Context, programare *models.Programare) (int, error) {
	query := fmt.Sprintf(
		"UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?",
		utils.AppointmentTableName,
		utils.ColumnIDPacient,
		utils.ColumnIDDoctor,
		utils.ColumnDate,
		utils.ColumnStatus,
		utils.ColumnIDProgramare,
	)

	res, err := db.ExecContext(ctx, query, programare.IDPacient, programare.IDDoctor, programare.Date, programare.Status, programare.IDProgramare)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to update programare: %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("[APPOINTMENT] Error getting rows affected: %v", err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[APPOINTMENT] No appointment found with ID %d to update.", programare.IDDoctor)
	} else {
		log.Printf("[APPOINTMENT] Successfully updated appointment with ID %d.", programare.IDDoctor)
	}

	return int(rowsAffected), nil
}

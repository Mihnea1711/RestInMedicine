package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) UpdateAppointmentByID(ctx context.Context, appointment *models.Appointment) (int, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf(
		"UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?",
		utils.AppointmentTableName,
		utils.ColumnIDPacient,
		utils.ColumnIDDoctor,
		utils.ColumnDate,
		utils.ColumnStatus,
		utils.ColumnIDProgramare,
	)

	log.Println("[APPOINTMENT] Attempting to update appointment")

	// Execute the SQL statement
	res, err := db.ExecContext(ctx, query, appointment.IDPacient, appointment.IDDoctor, appointment.Date, appointment.Status, appointment.IDProgramare)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to update appointment: %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("[APPOINTMENT] Error getting rows affected: %v", err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[APPOINTMENT] No appointment found with ID %d to update.", appointment.IDDoctor)
	} else {
		log.Printf("[APPOINTMENT] Successfully updated appointment with ID %d.", appointment.IDDoctor)
	}

	return int(rowsAffected), nil
}

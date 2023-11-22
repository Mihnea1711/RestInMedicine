package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) SaveAppointment(ctx context.Context, appointment *models.Appointment) (int, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf(
		"INSERT INTO %s (%s, %s, %s, %s) VALUES (?, ?, ?, ?)",
		utils.AppointmentTableName,
		utils.ColumnIDPacient,
		utils.ColumnIDDoctor,
		utils.ColumnDate,
		utils.ColumnStatus,
	)

	log.Println("[APPOINTMENT] Attempting to save appointment")

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, appointment.IDPacient, appointment.IDDoctor, appointment.Date, appointment.Status)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to save appointment: %v", err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[APPOINTMENT] Error getting last insert id whe saving appointment: %v", err)
		return 0, err
	}

	if lastInsertID == 0 {
		log.Printf("[APPOINTMENT] Something unexpected happened and the appointment could not be saved.")
	} else {
		log.Printf("[APPOINTMENT] Appointment saved successfully. ID: %d", lastInsertID)
	}

	return int(lastInsertID), nil
}

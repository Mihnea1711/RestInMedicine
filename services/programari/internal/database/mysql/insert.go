package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) SaveProgramare(ctx context.Context, programare *models.Programare) (int, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (%s, %s, %s, %s) VALUES (?, ?, ?, ?)",
		utils.AppointmentTableName,
		utils.ColumnIDPacient,
		utils.ColumnIDDoctor,
		utils.ColumnDate,
		utils.ColumnStatus,
	)

	result, err := db.ExecContext(ctx, query, programare.IDPacient, programare.IDDoctor, programare.Date, programare.Status)
	if err != nil {
		// Check if the error is due to a duplicate entry violation
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			log.Printf("[APPOINTMENT] Conflict error executing query to save programare: %v", err)
			return 0, nil
		}

		log.Printf("[APPOINTMENT] Error executing query to save programare: %v", err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[APPOINTMENT] Error getting last insert id whe saving appointment: %v", err)
		return 0, err
	}

	log.Println("[APPOINTMENT] Appointment saved successfully.")
	return int(lastInsertID), nil
}

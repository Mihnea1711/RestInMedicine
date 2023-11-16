package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) SaveProgramare(ctx context.Context, programare *models.Programare) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (%s, %s, %s, %s) VALUES (?, ?, ?, ?)",
		utils.AppointmentTableName,
		utils.ColumnIDPacient,
		utils.ColumnIDDoctor,
		utils.ColumnDate,
		utils.ColumnStatus,
	)

	_, err := db.ExecContext(ctx, query, programare.IDPacient, programare.IDDoctor, programare.Date, programare.Status)
	if err != nil {
		log.Printf("[APPOINTMENT] Error executing query to save programare: %v", err)
		return err
	}

	log.Println("[APPOINTMENT] Programare saved successfully.")
	return nil
}

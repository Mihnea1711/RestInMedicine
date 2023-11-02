package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (db *MySQLDatabase) UpdateProgramareByID(ctx context.Context, programare *models.Programare) (int64, error) {
	query := fmt.Sprintf(`UPDATE %s SET id_pacient = ?, id_doctor = ?, date = ?, status = ? WHERE id_programare = ?`, utils.PROGRAMARE_TABLE)

	res, err := db.ExecContext(ctx, query, programare.IDPacient, programare.IDDoctor, programare.Date, programare.Status, programare.IDProgramare)
	if err != nil {
		log.Printf("[PROGRAMARE] Error executing query to update programare: %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("[PROGRAMARE] Error getting rows affected: %v", err)
		return 0, err
	}

	log.Printf("[PROGRAMARE] %d programare updated successfully.", rowsAffected)
	return rowsAffected, nil
}

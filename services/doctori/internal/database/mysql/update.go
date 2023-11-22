package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) UpdateDoctorByID(ctx context.Context, doctor *models.Doctor) (int, error) {
	// Construct the SQL update query
	query := fmt.Sprintf(`
	UPDATE %s 
	SET %s = ?, %s = ?, %s = ?, %s = ?, %s = ? 
	WHERE %s = ?`,
		utils.DoctorTableName,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnSpecializare,
		utils.ColumnIDDoctor,
	)

	log.Printf("[DOCTOR] Attempting to update doctor with ID %d...", doctor.IDDoctor)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, doctor.Nume, doctor.Prenume, doctor.Email, doctor.Telefon, doctor.Specializare, doctor.IDDoctor)
	if err != nil {
		log.Printf("[DOCTOR] Error executing query to update doctor with ID %d: %v", doctor.IDDoctor, err)
		return 0, err
	}

	// Get the number of rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[DOCTOR] Error fetching rows affected for update query of doctor with ID %d: %v", doctor.IDDoctor, err)
		return 0, err
	}

	if rowsAffected == 0 {
		log.Printf("[DOCTOR] No doctor found with ID %d to update.", doctor.IDDoctor)
	} else {
		log.Printf("[DOCTOR] Successfully updated doctor with ID %d. Rows affected: %d", doctor.IDDoctor, rowsAffected)
	}

	return int(rowsAffected), nil
}

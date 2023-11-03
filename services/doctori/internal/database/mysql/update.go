package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) UpdateDoctorByID(ctx context.Context, doctor *models.Doctor) (int64, error) {
	// Construct the SQL update query
	query := fmt.Sprintf(`
		UPDATE %s 
		SET nume = ?, prenume = ?, email = ?, telefon = ?, specializare = ? 
		WHERE id_doctor = ?`, utils.DOCTOR_TABLE)

	log.Printf("[DOCTOR] Executing update query for doctor with ID %d...", doctor.IDDoctor)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, doctor.Nume, doctor.Prenume, doctor.Email, doctor.Telefon, doctor.Specializare, doctor.IDDoctor)
	if err != nil {
		log.Printf("[DOCTOR] Error updating doctor with ID %d: %v", doctor.IDDoctor, err)
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
		log.Printf("[DOCTOR] Successfully updated doctor with ID %d.", doctor.IDDoctor)
	}

	return rowsAffected, nil
}

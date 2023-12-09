package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) SaveDoctor(ctx context.Context, doctor *models.Doctor) (int, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s, %s) VALUES (?, ?, ?, ?, ?, ?, ?)",
		utils.DoctorTableName,
		utils.ColumnIDUser,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnSpecialization,
		utils.ColumnIsActive,
	)

	log.Println("[PATIENT] Attempting to save doctor")

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, doctor.IDUser, doctor.FirstName, doctor.SecondName, doctor.Email, doctor.PhoneNumber, doctor.Specialization, doctor.IsActive)
	if err != nil {
		log.Printf("[DOCTOR] Error executing query to save doctor: %v", err)
		return 0, fmt.Errorf("failed to save doctor: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[DOCTOR] Error getting last insert id while saving doctor: %v", err)
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	if lastInsertID == 0 {
		log.Printf("[DOCTOR] Something unexpected happened and the doctor could not be saved.")
	} else {
		log.Printf("[DOCTOR] Doctor saved successfully. ID: %d", lastInsertID)
	}

	return int(lastInsertID), nil
}

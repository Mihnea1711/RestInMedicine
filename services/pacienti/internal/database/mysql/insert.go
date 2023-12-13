package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) SavePatient(ctx context.Context, patient *models.Patient) (int, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s, %s, %s) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		utils.PatientTableName,
		utils.ColumnIDUser,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnCNP,
		utils.ColumnBirthDay,
		utils.ColumnIsActive,
	)

	log.Println("[PATIENT] Attempting to save patient")

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, patient.IDUser, patient.FirstName, patient.SecondName, patient.Email, patient.PhoneNumber, patient.CNP, patient.BirthDay, patient.IsActive)
	if err != nil {
		log.Printf("[PATIENT] Error executing query to save patient: %v", err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[PATIENT] Error getting last insert id when saving patient: %v", err)
		return 0, err
	}

	if lastInsertID == 0 {
		log.Printf("[PATIENT] Something unexpected happened and the patient could not be saved.")
	} else {
		log.Printf("[PATIENT] Patient saved successfully. ID: %d", lastInsertID)
	}

	return int(lastInsertID), nil
}

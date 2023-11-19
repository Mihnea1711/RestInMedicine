package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) SaveDoctor(ctx context.Context, doctor *models.Doctor) (int, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s) VALUES (?, ?, ?, ?, ?, ?)",
		utils.DoctorTableName,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnSpecializare,
	)

	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, doctor.IDUser, doctor.Nume, doctor.Prenume, doctor.Email, doctor.Telefon, doctor.Specializare)
	if err != nil {
		// Check if the error is due to a duplicate entry violation
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			return 0, nil
		}

		log.Printf("[DOCTOR] Error executing query to save doctor: %v", err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[PATIENT] Error getting last insert id whe saving patient: %v", err)
		return 0, err
	}

	log.Println("[DOCTOR] Doctor saved successfully.")
	return int(lastInsertID), nil
}

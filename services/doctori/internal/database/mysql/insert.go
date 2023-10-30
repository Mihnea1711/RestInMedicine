package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) SaveDoctor(ctx context.Context, doctor *models.Doctor) error {
	// Construct the SQL insert query
	query := fmt.Sprintf(`INSERT INTO %s (id_user, nume, prenume, email, telefon, specializare) VALUES (?, ?, ?, ?, ?, ?)`, utils.DOCTOR_TABLE)

	// Execute the SQL statement
	_, err := db.ExecContext(ctx, query, doctor.IDUser, doctor.Nume, doctor.Prenume, doctor.Email, doctor.Telefon, doctor.Specializare)
	if err != nil {
		log.Printf("Error executing query to save doctor: %v", err)
		return err
	}

	log.Println("Doctor saved successfully.")
	return nil
}

package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (db *MySQLDatabase) SavePacient(ctx context.Context, pacient *models.Pacient) (int, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s, %s, %s) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		utils.TableName,
		utils.ColumnIDUser,
		utils.ColumnNume,
		utils.ColumnPrenume,
		utils.ColumnEmail,
		utils.ColumnTelefon,
		utils.ColumnCNP,
		utils.ColumnDataNasterii,
		utils.ColumnIsActive,
	)
	// Execute the SQL statement
	result, err := db.ExecContext(ctx, query, pacient.IDUser, pacient.Nume, pacient.Prenume, pacient.Email, pacient.Telefon, pacient.CNP, pacient.DataNasterii, pacient.IsActive)
	if err != nil {
		// Check if the error is due to a duplicate entry violation
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			return 0, nil
		}

		log.Printf("[PATIENT] Error executing query to save pacient: %v", err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[PATIENT] Error getting last insert id whe saving patient: %v", err)
		return 0, err
	}

	log.Println("[PATIENT] Pacient saved successfully.")
	return int(lastInsertID), nil
}

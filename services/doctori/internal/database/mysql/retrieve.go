package mysql

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (db *MySQLDatabase) FetchDoctors(ctx context.Context, filters map[string]interface{}, page, limit int) ([]models.Doctor, error) {
	// Get the offset based on page and limit
	offset := (page - 1) * limit

	qb := squirrel.Select("*").From(utils.DoctorTableName)

	// Check if filters is not empty and add WHERE clause if needed
	if len(filters) > 0 {
		// Check if filters contain 'first_name'
		if firstName, ok := filters[utils.ColumnFirstName].(string); ok && firstName != "" {
			// Apply the LIKE filter for the 'first_name' column
			qb = qb.Where(squirrel.Like{utils.ColumnFirstName: "%" + strings.ToLower(firstName) + "%"})
		}

		// Check if filters contain 'isActive'
		if isActive, ok := filters[utils.ColumnIsActive].(bool); ok {
			// Apply the filter for 'isActive'
			qb = qb.Where(squirrel.Eq{utils.ColumnIsActive: isActive})
		}

		// Check if filters contain 'specialization'
		if specialization, ok := filters[utils.ColumnSpecialization].(models.Specialization); ok && specialization != "" {
			// Apply the filter for 'specialization'
			qb = qb.Where(squirrel.Eq{utils.ColumnSpecialization: specialization})
		}
	}

	// Add LIMIT and OFFSET for pagination
	qb = qb.Limit(uint64(limit)).Offset(uint64(offset))

	query, args, err := qb.ToSql()
	if err != nil {
		log.Printf("[DOCTOR] FetchDoctors: Failed to construct SQL query: %v", err)
		return nil, fmt.Errorf("internal server error")
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("[DOCTOR] FetchDoctors: Failed to query database: %v", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer rows.Close()

	var doctors []models.Doctor
	for rows.Next() {
		var doctor models.Doctor
		err := rows.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.FirstName, &doctor.SecondName, &doctor.Email, &doctor.PhoneNumber, &doctor.Specialization, &doctor.IsActive)
		if err != nil {
			log.Printf("[DOCTOR] Error scanning doctor row: %v", err)
			return nil, err
		}
		doctors = append(doctors, doctor)
	}

	// Check for errors from iterating over rows.
	err = rows.Err()
	if err != nil {
		log.Printf("[DOCTOR] Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched %d doctors.", len(doctors))
	return doctors, nil
}

func (db *MySQLDatabase) FetchDoctorByID(ctx context.Context, doctorID int) (*models.Doctor, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDDoctor,
		utils.ColumnIDUser,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnSpecialization,
		utils.ColumnIsActive,
		utils.DoctorTableName,
		utils.ColumnIDDoctor,
	)

	log.Printf("[PATIENT] Attempting to fetch doctor with ID %d", doctorID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, doctorID)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.FirstName, &doctor.SecondName, &doctor.Email, &doctor.PhoneNumber, &doctor.Specialization, &doctor.IsActive)
	if err != nil {
		log.Printf("[DOCTOR] Error fetching doctor by ID %d: %v", doctorID, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by ID %d.", doctorID)
	return &doctor, nil
}

func (db *MySQLDatabase) FetchDoctorByEmail(ctx context.Context, email string) (*models.Doctor, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDDoctor,
		utils.ColumnIDUser,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnSpecialization,
		utils.ColumnIsActive,
		utils.DoctorTableName,
		utils.ColumnEmail,
	)

	log.Printf("[PATIENT] Attempting to fetch doctor with email %s", email)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, email)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.FirstName, &doctor.SecondName, &doctor.Email, &doctor.PhoneNumber, &doctor.Specialization, &doctor.IsActive)
	if err != nil {
		log.Printf("[DOCTOR] Error fetching doctor by email %s: %v", email, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by email %s.", email)
	return &doctor, nil
}

func (db *MySQLDatabase) FetchDoctorByUserID(ctx context.Context, userID int) (*models.Doctor, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDDoctor,
		utils.ColumnIDUser,
		utils.ColumnFirstName,
		utils.ColumnSecondName,
		utils.ColumnEmail,
		utils.ColumnPhoneNumber,
		utils.ColumnSpecialization,
		utils.ColumnIsActive,
		utils.DoctorTableName,
		utils.ColumnIDUser,
	)

	log.Printf("[PATIENT] Attempting to fetch doctor with user ID %d", userID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, userID)

	var doctor models.Doctor
	err := row.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.FirstName, &doctor.SecondName, &doctor.Email, &doctor.PhoneNumber, &doctor.Specialization, &doctor.IsActive)
	if err != nil {
		log.Printf("[DOCTOR] Error fetching doctor by user ID %d: %v", userID, err)
		return nil, err
	}

	log.Printf("[DOCTOR] Successfully fetched doctor by user ID %d.", userID)
	return &doctor, nil
}

package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/config"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

type MySQLDatabase struct {
	*sql.DB
}

func NewMySQL(config *config.MySQLConfig) (*MySQLDatabase, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
		config.Charset,
		config.ParseTime,
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Printf("Error connecting to MySQL: %v", err)
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime * time.Second)

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Printf("Error pinging MySQL: %v", err)
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	return &MySQLDatabase{DB: db}, nil
}

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

func (db *MySQLDatabase) FetchDoctors() ([]models.Doctor, error) {
	query := fmt.Sprintf(`SELECT id_doctor, id_user, nume, prenume, email, telefon, specializare FROM %s`, utils.DOCTOR_TABLE)

	// Execute the SQL query
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query to fetch doctors: %v", err)
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor

	for rows.Next() {
		var doctor models.Doctor
		err := rows.Scan(&doctor.IDDoctor, &doctor.IDUser, &doctor.Nume, &doctor.Prenume, &doctor.Email, &doctor.Telefon, &doctor.Specializare)
		if err != nil {
			log.Printf("Error scanning doctor row: %v", err)
			return nil, err
		}
		doctors = append(doctors, doctor)
	}

	// Check for errors from iterating over rows.
	err = rows.Err()
	if err != nil {
		log.Printf("Error after iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("Successfully fetched %d doctors.", len(doctors))
	return doctors, nil
}

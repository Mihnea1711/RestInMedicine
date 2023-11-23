package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/mihnea1711/POS_Project/services/idm/internal/database"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/config"
)

type MySQLDatabase struct {
	*sql.DB
}

// NewMySQL creates a new MySQL database connection.
func NewMySQL(ctx context.Context, config *config.MySQLConfig) (database.Database, error) {
	// Construct the MySQL connection string
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

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Printf("[IDM] Error connecting to MySQL: %v", err)
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	// Set connection pool parameters
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime * time.Second)

	// Test the connection by pinging the database
	if err := db.PingContext(ctx); err != nil {
		log.Printf("[IDM] Error pinging MySQL: %v", err)
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	// Log successful connection
	log.Println("[IDM] Successfully connected to MySQL database")

	// Return the MySQLDatabase instance
	return &MySQLDatabase{DB: db}, nil
}

// GetDB returns the underlying SQL database instance.
func (db *MySQLDatabase) GetDB() *sql.DB {
	return db.DB
}

// Close closes the MySQL database connection.
func (db *MySQLDatabase) Close() error {
	// Log the attempt to close the MySQL database connection
	log.Println("[IDM] Closing MySQL database connection")

	// Close the connection and return any error that occurs
	return db.DB.Close()
}

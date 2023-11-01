package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/database"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/config"
)

type MySQLDatabase struct {
	*sql.DB
}

func NewMySQL(config *config.MySQLConfig) (database.Database, error) {
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

	fmt.Println(connStr)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Printf("[PACIENTI] Error connecting to MySQL: %v", err)
		return nil, fmt.Errorf("[PACIENTI] Failed to connect to MySQL: %v", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime * time.Second)

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Printf("[PACIENTI] Error pinging MySQL: %v", err)
		return nil, fmt.Errorf("[PACIENTI] Failed to ping MySQL: %v", err)
	}

	return &MySQLDatabase{DB: db}, nil
}

func (db *MySQLDatabase) Close() error {
	return db.DB.Close()
}

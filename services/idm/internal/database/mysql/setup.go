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

func NewMySQL(ctx context.Context, config *config.MySQLConfig) (database.Database, error) {
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
		log.Printf("[IDM] Error connecting to MySQL: %v", err)
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime * time.Second)

	// Test the connection
	if err := db.PingContext(ctx); err != nil {
		log.Printf("[IDM] Error pinging MySQL: %v", err)
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	return &MySQLDatabase{DB: db}, nil
}

func (db *MySQLDatabase) Close() error {
	return db.DB.Close()
}

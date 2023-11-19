package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// AddUserToDB adds a user to the MySQL database and returns the new user's ID.
func (db *MySQLDatabase) AddUser(ctx context.Context, newUser models.UserRegistration) (int, error) {
	// Start a database transaction
	tx, err := db.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		errMsg := fmt.Sprintf("Error starting a database transaction: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Insert the user into the User table
	userQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)",
		utils.UserTable,
		utils.ColumnUserName,
		utils.ColumnUserPassword,
	)
	result, err := tx.ExecContext(ctx, userQuery, newUser.Username, newUser.Password)
	if err != nil {
		tx.Rollback()

		// Check if the error is due to a duplicate entry violation
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			// Return a conflict error status
			return 0, nil
		}

		errMsg := fmt.Sprintf("Error adding user to User table: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Get the new user's ID
	userID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Printf("[IDM] Error getting last inserted ID from User Table: %v", err)
		return 0, fmt.Errorf("error getting last inserted ID from User Table: %v", err)
	}

	if userID == 0 {
		tx.Rollback()
		log.Printf("[IDM] Last inserted ID is 0 in User Table. No rows were affected.")
		return 0, nil
	}

	// Insert the user's role into the Role table
	roleQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)",
		utils.RoleTable,
		utils.ColumnRoleIDUser,
		utils.ColumnRole,
	)

	result, err = tx.ExecContext(ctx, roleQuery, userID, newUser.Role)
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Sprintf("Error adding user's role to Role table: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Get the new user's ID
	userRoleID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Printf("[IDM] Error getting last inserted ID from Role Table: %v", err)
		return 0, fmt.Errorf("error getting last inserted ID from Role Table: %v", err)
	}

	if userRoleID == 0 {
		tx.Rollback()
		log.Printf("[IDM] Last inserted ID is 0 in Role Table. No rows were affected.")
		return 0, nil
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		errMsg := fmt.Sprintf("Error committing the transaction: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	log.Printf("[IDM] Added user to DB: User and Role tables updated")
	return int(userID), nil
}

// GetAllUsersFromDB retrieves all users from the MySQL database.
func (db *MySQLDatabase) GetAllUsers(ctx context.Context, page, limit int) ([]models.User, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT * FROM %s LIMIT ? OFFSET ?",
		utils.UserTable,
	)
	rows, err := db.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		errMsg := fmt.Sprintf("Error retrieving all users from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return nil, fmt.Errorf(errMsg)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.IDUser, &user.Username, &user.Password)
		if err != nil {
			errMsg := fmt.Sprintf("Error scanning user data from DB: %v", err)
			log.Printf("[IDM] %s", errMsg)
			return nil, fmt.Errorf(errMsg)
		}
		users = append(users, user)
	}
	log.Printf("[IDM] Retrieved all users from DB")
	return users, nil
}

// GetUserFromDBByID retrieves a user from the MySQL database by user ID.
func (db *MySQLDatabase) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnIDUser,
	)
	var user models.User
	err := db.DB.QueryRowContext(ctx, query, userID).Scan(&user.IDUser, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return nil, nil
		}
		errMsg := fmt.Sprintf("Error retrieving user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return nil, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Retrieved user with ID %d from DB", userID)
	return &user, nil
}

// GetUserFromDBByUsername retrieves a user from the MySQL database by username.
func (db *MySQLDatabase) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserName,
	)
	var user models.User
	err := db.DB.QueryRowContext(ctx, query, username).Scan(&user.IDUser, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return nil, nil
		}
		errMsg := fmt.Sprintf("Error retrieving user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return nil, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Retrieved user with Username %s from DB", username)
	return &user, nil
}

// UpdateUserInDB updates a user in the MySQL database.
func (db *MySQLDatabase) UpdateUserByID(ctx context.Context, userCredentials models.CredentialsRequest, userID int) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserName,
		utils.ColumnIDUser,
	)
	result, err := db.DB.ExecContext(ctx, query, userCredentials.Username, userID)
	if err != nil {
		// Check if the error is due to a duplicate entry violation
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			// Return a conflict error status
			return 0, nil
		}

		errMsg := fmt.Sprintf("Error updating user in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[IDM] Error getting rows affected in User Table: %v", err)
		return 0, fmt.Errorf("error getting rows affected in User Table: %v", err)
	}
	if rowsAffected == 0 {
		// No rows were affected, indicating that the user was not found
		log.Printf("[IDM] User with ID %d not found", userID)
		return 0, nil
	}

	log.Printf("[IDM] Updated user in DB: %d rows affected", rowsAffected)
	return int(rowsAffected), nil
}

// DeleteUserFromDBByID deletes a user from the MySQL database by user ID.
func (db *MySQLDatabase) DeleteUserByID(ctx context.Context, userID int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnIDUser,
	)
	result, err := db.DB.ExecContext(ctx, query, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Error deleting user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[IDM] Error getting rows affected in User Table: %v", err)
		return 0, fmt.Errorf("error getting rows affected in User Table: %v", err)
	}
	if rowsAffected == 0 {
		// No rows were affected, indicating that the user was not found
		log.Printf("[IDM] User with ID %d not found", userID)
		return 0, nil
	}

	log.Printf("[IDM] Deleted user from DB: %d rows affected", rowsAffected)
	return int(rowsAffected), nil
}

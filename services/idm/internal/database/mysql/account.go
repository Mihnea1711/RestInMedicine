package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// GetUserPasswordByUsername retrieves the hashed password of a user from the MySQL database by username.
func (db *MySQLDatabase) GetUserPasswordByUsername(ctx context.Context, username string) (string, error) {
	// Construct the SQL query to retrieve the user's hashed password by username
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", utils.ColumnUserPassword, utils.UserTable, utils.ColumnUserName)

	// Variable to store the retrieved password
	var password string

	// Execute the SQL query to retrieve the user's hashed password
	err := db.DB.QueryRowContext(ctx, query, username).Scan(&password)
	if err != nil {
		// Check if the user was not found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return "", err
		}

		// Log and return an error if the retrieval fails
		errMsg := fmt.Sprintf("Error retrieving user's password from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return "", err
	}

	// Log that the password has been retrieved successfully
	log.Printf("[IDM] Retrieved password for user with Username %s from DB", username)
	return password, nil
}

// GetUserRoleByUserID retrieves the role of a user from the MySQL database by user ID.
func (db *MySQLDatabase) GetUserRoleByUserID(ctx context.Context, userID int) (string, error) {
	// Construct the SQL query to retrieve the user's role by user ID
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", utils.ColumnRole, utils.RoleTable, utils.ColumnRoleIDUser)

	// Variable to store the retrieved role
	var role string

	// Execute the SQL query to retrieve the user's role
	err := db.DB.QueryRowContext(ctx, query, userID).Scan(&role)
	if err != nil {
		// Log and return an error if the retrieval fails
		errMsg := fmt.Sprintf("Error retrieving user's role from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return "", err
	}

	// Log that the role has been retrieved successfully
	log.Printf("[IDM] Retrieved role for user with ID %d from DB", userID)
	return role, nil
}

// GetUserRoleByUsername retrieves the role of a user from the MySQL database by username.
func (db *MySQLDatabase) GetUserRoleByUsername(ctx context.Context, username string) (string, error) {
	// Construct the SQL query to retrieve the user's role by username
	query := fmt.Sprintf("SELECT %s FROM %s %s JOIN %s %s ON %s.%s = %s.%s WHERE %s.%s = ?",
		utils.ColumnRole,
		utils.RoleTable,
		utils.AliasRole,
		utils.UserTable,
		utils.AliasUser,
		utils.AliasRole,
		utils.ColumnRoleIDUser,
		utils.AliasUser,
		utils.ColumnIDUser,
		utils.AliasUser,
		utils.ColumnUserName,
	)

	// Variable to store the retrieved role
	var role string

	// Execute the SQL query to retrieve the user's role
	err := db.DB.QueryRowContext(ctx, query, username).Scan(&role)
	if err != nil {
		// Log and return an error if the retrieval fails
		errMsg := fmt.Sprintf("Error retrieving user's role from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return "", err
	}

	// Log that the role has been retrieved successfully
	log.Printf("[IDM] Retrieved role for user with Username %s from DB", username)
	return role, nil
}

// UpdateUserRoleByUserID updates a user's role in the MySQL database by user ID and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserRoleByUserID(ctx context.Context, userID int, newRole string) (int, error) {
	// Construct the SQL query to update the user's role by user ID
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.RoleTable,
		utils.ColumnRole,
		utils.ColumnRoleIDUser,
	)

	// Execute the SQL query to update the user's role
	result, err := db.DB.ExecContext(ctx, query, newRole, userID)
	if err != nil {
		// Log and return an error if the update fails
		errMsg := fmt.Sprintf("Error changing user's role in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Get the number of rows affected by the update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Log and return an error if getting the affected rows count fails
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Check if no rows were affected, indicating that the user was not found
	if rowsAffected == 0 {
		msg := "No rows were changed"
		log.Printf("[IDM] %s", msg)
		return 0, nil
	}

	// Log that the role has been successfully updated
	log.Printf("[IDM] Changed role for user with ID %d to %s in DB", userID, newRole)
	return int(rowsAffected), nil
}

// UpdateUserRoleByUsername updates a user's role in the MySQL database by username and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserRoleByUsername(ctx context.Context, username string, newRole string) (int, error) {
	// Construct the SQL query to update the user's role by username
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = (SELECT %s FROM %s WHERE %s = ?)",
		utils.RoleTable,
		utils.ColumnRole,
		utils.ColumnRoleIDUser,
		utils.ColumnIDUser,
		utils.UserTable,
		utils.ColumnUserName,
	)

	// Execute the SQL query to update the user's role
	result, err := db.DB.ExecContext(ctx, query, newRole, username)
	if err != nil {
		// Log and return an error if the update fails
		errMsg := fmt.Sprintf("Error changing user's role in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Get the number of rows affected by the update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Log and return an error if getting the affected rows count fails
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Check if no rows were affected, indicating that the user was not found
	if rowsAffected == 0 {
		msg := "No rows were changed"
		log.Printf("[IDM] %s", msg)
		return 0, nil
	}

	// Log that the role has been successfully updated
	log.Printf("[IDM] Changed role for user with Username %s to %s in DB", username, newRole)
	return int(rowsAffected), nil
}

// UpdateUserPasswordByUserID updates a user's password in the MySQL database by user ID and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserPasswordByUserID(ctx context.Context, userID int, newPassword string) (int, error) {
	// Construct the SQL query to update the user's password by user ID
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserPassword,
		utils.ColumnIDUser,
	)

	// Execute the SQL query to update the user's password
	result, err := db.DB.ExecContext(ctx, query, newPassword, userID)
	if err != nil {
		// Log and return an error if the update fails
		errMsg := fmt.Sprintf("Error changing user's password in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Get the number of rows affected by the update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Log and return an error if getting the affected rows count fails
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Check if no rows were affected, indicating that the user was not found
	if rowsAffected == 0 {
		msg := "No rows were changed"
		log.Printf("[IDM] %s", msg)
		return 0, nil
	}

	// Log that the password has been successfully updated
	log.Printf("[IDM] Changed password for user with ID %d in DB", userID)
	return int(rowsAffected), nil
}

// UpdateUserPasswordByUsername updates a user's password in the MySQL database by username and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserPasswordByUsername(ctx context.Context, username string, newPassword string) (int, error) {
	// Construct the SQL query to update the user's password by username
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserPassword,
		utils.ColumnUserName,
	)

	// Execute the SQL query to update the user's password
	result, err := db.DB.ExecContext(ctx, query, newPassword, username)
	if err != nil {
		// Log and return an error if the update fails
		errMsg := fmt.Sprintf("Error changing user's password in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Get the number of rows affected by the update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Log and return an error if getting the affected rows count fails
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Check if no rows were affected, indicating that the user was not found
	if rowsAffected == 0 {
		msg := "No rows were changed"
		log.Printf("[IDM] %s", msg)
		return 0, nil
	}

	// Log that the password has been successfully updated
	log.Printf("[IDM] Changed password for user with Username %s in DB", username)
	return int(rowsAffected), nil
}

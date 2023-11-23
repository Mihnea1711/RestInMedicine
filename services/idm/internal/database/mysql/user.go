package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

	log.Println("[IDM] Attempting to save user")

	result, err := tx.ExecContext(ctx, userQuery, newUser.Username, newUser.Password)
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Sprintf("Error adding user to User table: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, err
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
		return 0, fmt.Errorf("an unexpected error happened while saving user: %v", err)
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
		return 0, fmt.Errorf("an unexpected error happened while saving user's role: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		errMsg := fmt.Sprintf("Error committing the transaction: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	log.Printf("[IDM] Added user to DB: User and Role tables updated. User ID: %d", userID)
	return int(userID), nil
}

// GetAllUsersFromDB retrieves all users from the MySQL database.
func (db *MySQLDatabase) GetAllUsers(ctx context.Context, page, limit int) ([]models.User, error) {
	// Calculate the offset based on the page and limit
	offset := (page - 1) * limit

	// Construct the SQL query to retrieve users with pagination
	query := fmt.Sprintf("SELECT * FROM %s LIMIT ? OFFSET ?",
		utils.UserTable,
	)

	log.Printf("[IDM] Attempting to fetch users with limit=%d, offset=%d", limit, offset)

	// Execute the SQL query
	rows, err := db.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		errMsg := fmt.Sprintf("Error retrieving all users from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return nil, fmt.Errorf(errMsg)
	}
	defer rows.Close()

	// Slice to store the retrieved users
	var users []models.User

	// Iterate through the result set
	for rows.Next() {
		var user models.User

		// Scan user data from the result set
		err := rows.Scan(&user.IDUser, &user.Username, &user.Password)
		if err != nil {
			errMsg := fmt.Sprintf("Error scanning user data from DB: %v", err)
			log.Printf("[IDM] %s", errMsg)
			return nil, fmt.Errorf(errMsg)
		}

		// Append the user to the slice
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("[IDM] Error after iterating over rows: %v", err)
		return nil, err
	}

	// Log that all users have been retrieved successfully
	log.Printf("[IDM] Retrieved %d users from DB", len(users))
	return users, nil
}

// GetUserByID retrieves a user from the MySQL database by user ID.
func (db *MySQLDatabase) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	// Construct the SQL query to retrieve a user by ID
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnIDUser,
	)

	log.Printf("[IDM] Attempting to fetch user with ID %d", userID)

	// Create a user model to store the retrieved user
	var user models.User

	// Execute the SQL query and scan the result into the user model
	err := db.DB.QueryRowContext(ctx, query, userID).Scan(&user.IDUser, &user.Username, &user.Password)
	if err != nil {
		// Log and return an error for other errors
		errMsg := fmt.Sprintf("Error retrieving user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return nil, err
	}

	// Log that the user has been retrieved successfully
	log.Printf("[IDM] Retrieved user with ID %d from DB", userID)
	return &user, nil
}

// GetUserByUsername retrieves a user from the MySQL database by username.
func (db *MySQLDatabase) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	// Construct the SQL query to retrieve a user by username
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserName,
	)

	// Create a user model to store the retrieved user
	var user models.User

	// Execute the SQL query and scan the result into the user model
	err := db.DB.QueryRowContext(ctx, query, username).Scan(&user.IDUser, &user.Username, &user.Password)
	if err != nil {
		// Log and return an error for other errors
		errMsg := fmt.Sprintf("Error retrieving user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return nil, err
	}

	// Log that the user has been retrieved successfully
	log.Printf("[IDM] Retrieved user with Username %s from DB", username)
	return &user, nil
}

// UpdateUserByID updates a user in the MySQL database.
func (db *MySQLDatabase) UpdateUserByID(ctx context.Context, userCredentials models.CredentialsRequest, userID int) (int, error) {
	// Construct the SQL query to update a user by ID
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserName,
		utils.ColumnIDUser,
	)

	// Execute the SQL query to update the user's username
	result, err := db.DB.ExecContext(ctx, query, userCredentials.Username, userID)
	if err != nil {
		// Log and return an error for other errors
		errMsg := fmt.Sprintf("Error updating user in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, err
	}

	// Get the number of rows affected by the update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[IDM] Error getting rows affected in User Table: %v", err)
		return 0, fmt.Errorf("error getting rows affected in User Table: %v", err)
	}

	// Check if no rows were affected, indicating that the user was not found
	if rowsAffected == 0 {
		log.Printf("[IDM] User with ID %d not found", userID)
		return 0, nil
	}

	// Log that the user has been updated successfully
	log.Printf("[IDM] Updated user in DB: %d rows affected", rowsAffected)
	return int(rowsAffected), nil
}

// DeleteUserByID deletes a user from the MySQL database by user ID.
func (db *MySQLDatabase) DeleteUserByID(ctx context.Context, userID int) (int, error) {
	// Construct the SQL query to delete a user by ID
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnIDUser,
	)

	// Execute the SQL query to delete the user
	result, err := db.DB.ExecContext(ctx, query, userID)
	if err != nil {
		// Log and return an error if the deletion fails
		errMsg := fmt.Sprintf("Error deleting user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, err
	}

	// Get the number of rows affected by the deletion
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[IDM] Error getting rows affected in User Table: %v", err)
		return 0, fmt.Errorf("error getting rows affected in User Table: %v", err)
	}

	// Check if no rows were affected, indicating that the user was not found
	if rowsAffected == 0 {
		log.Printf("[IDM] User with ID %d not found", userID)
		return 0, nil
	}

	// Log that the user has been deleted successfully
	log.Printf("[IDM] Deleted user from DB: %d rows affected", rowsAffected)
	return int(rowsAffected), nil
}

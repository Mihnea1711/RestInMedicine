package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// AddUserToDB adds a user to the MySQL database and returns the new user's ID.
func (db *MySQLDatabase) AddUser(newUser models.UserRegistration) (int, error) {
	// Start a database transaction
	tx, err := db.DB.Begin()
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
	result, err := tx.Exec(userQuery, newUser.Username, newUser.Password)
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Sprintf("Error adding user to User table: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	// Get the new user's ID
	userID, _ := result.LastInsertId()

	// Insert the user's role into the Role table
	roleQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)",
		utils.RoleTable,
		utils.ColumnRoleIDUser,
		utils.ColumnRole,
	)
	_, err = tx.Exec(roleQuery, userID, newUser.Role)
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Sprintf("Error adding user's role to Role table: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
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
func (db *MySQLDatabase) GetAllUsers(page, limit int) ([]models.User, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT * FROM %s LIMIT ? OFFSET ?",
		utils.UserTable,
	)
	rows, err := db.DB.Query(query, limit, offset)
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
func (db *MySQLDatabase) GetUserByID(userID int) (models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnIDUser,
	)
	var user models.User
	err := db.DB.QueryRow(query, userID).Scan(&user.IDUser, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return models.User{}, fmt.Errorf("user not found")
		}
		errMsg := fmt.Sprintf("Error retrieving user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return models.User{}, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Retrieved user with ID %d from DB", userID)
	return user, nil
}

// GetUserFromDBByUsername retrieves a user from the MySQL database by username.
func (db *MySQLDatabase) GetUserByUsername(username string) (models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserName,
	)
	var user models.User
	err := db.DB.QueryRow(query, username).Scan(&user.IDUser, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return models.User{}, fmt.Errorf("user not found")
		}
		errMsg := fmt.Sprintf("Error retrieving user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return models.User{}, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Retrieved user with Username %s from DB", username)
	return user, nil
}

// UpdateUserInDB updates a user in the MySQL database.
func (db *MySQLDatabase) UpdateUserByID(userCredentials models.CredentialsRequest, userId int) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserName,
		utils.ColumnIDUser,
	)
	result, err := db.DB.Exec(query, userCredentials.Username, userId)
	if err != nil {
		errMsg := fmt.Sprintf("Error updating user in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("[IDM] Updated user in DB: %d rows affected", rowsAffected)
	return int(rowsAffected), nil
}

// DeleteUserFromDBByID deletes a user from the MySQL database by user ID.
func (db *MySQLDatabase) DeleteUserByID(userID int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?",
		utils.UserTable,
		utils.ColumnIDUser,
	)
	result, err := db.DB.Exec(query, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Error deleting user from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("[IDM] Deleted user from DB: %d rows affected", rowsAffected)
	return int(rowsAffected), nil
}

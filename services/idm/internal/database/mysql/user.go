package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
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
	userQuery := "INSERT INTO User (Username, Password) VALUES (?, ?)"
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
	roleQuery := "INSERT INTO Role (IDUser, Role) VALUES (?, ?)"
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
func (db *MySQLDatabase) GetAllUsers() ([]models.User, error) {
	query := "SELECT * FROM User"
	rows, err := db.DB.Query(query)
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
	query := "SELECT * FROM User WHERE IDUser = ?"
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
	query := "SELECT * FROM User WHERE Username = ?"
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
	query := "UPDATE User SET Username = ?, Password = ? WHERE IDUser = ?"
	result, err := db.DB.Exec(query, userCredentials.Username, userCredentials.Password, userId)
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
	query := "DELETE FROM User WHERE IDUser = ?"
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

package mysql

import (
	"database/sql"
	"fmt"
	"log"
)

// GetUserPasswordByUsername retrieves the hashed password of a user from the MySQL database by username.
func (db *MySQLDatabase) GetUserPasswordByUsername(username string) (string, error) {
	query := "SELECT Password FROM User WHERE Username = ?"
	var password string
	err := db.DB.QueryRow(query, username).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return "", fmt.Errorf("user not found")
		}
		errMsg := fmt.Sprintf("Error retrieving user's password from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return "", fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Retrieved password for user with Username %s from DB", username)
	return password, nil
}

// GetUserRoleByUserID retrieves the role of a user from the MySQL database by user ID.
func (db *MySQLDatabase) GetUserRoleByUserID(userID int) (string, error) {
	query := "SELECT Role FROM Role WHERE IDUser = ?"
	var role string
	err := db.DB.QueryRow(query, userID).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return "", fmt.Errorf("user not found")
		}
		errMsg := fmt.Sprintf("Error retrieving user's role from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return "", fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Retrieved role for user with ID %d from DB", userID)
	return role, nil
}

// GetUserRoleByUsername retrieves the role of a user from the MySQL database by username.
func (db *MySQLDatabase) GetUserRoleByUsername(username string) (string, error) {
	query := "SELECT Role FROM Role r JOIN User u ON r.IDUser = u.IDUser WHERE u.Username = ?"
	var role string
	err := db.DB.QueryRow(query, username).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return "", fmt.Errorf("user not found")
		}
		errMsg := fmt.Sprintf("Error retrieving user's role from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return "", fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Retrieved role for user with Username %s from DB", username)
	return role, nil
}

// GetUserTokenByID retrieves the token of a user by their ID.
func (db *MySQLDatabase) GetUserTokenByID(userID int) (string, error) {
	query := "SELECT Token FROM User WHERE IDUser = ?"
	var userToken string
	err := db.DB.QueryRow(query, userID).Scan(&userToken)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return "", fmt.Errorf("user not found")
		}
		errMsg := fmt.Sprintf("Error retrieving user token from DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return "", fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Retrieved token for user with ID %d from DB", userID)
	return userToken, nil
}

// ChangeUserRoleByUserID updates a user's role in the MySQL database by user ID and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserRoleByUserID(userID int, newRole string) (int, error) {
	query := "UPDATE Role SET Role = ? WHERE IDUser = ?"
	result, err := db.DB.Exec(query, newRole, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Error changing user's role in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Changed role for user with ID %d to %s in DB", userID, newRole)
	return int(rowsAffected), nil
}

// ChangeUserRoleByUsername updates a user's role in the MySQL database by username and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserRoleByUsername(username string, newRole string) (int, error) {
	query := "UPDATE Role SET Role = ? WHERE IDUser = (SELECT IDUser FROM User WHERE Username = ?)"
	result, err := db.DB.Exec(query, newRole, username)
	if err != nil {
		errMsg := fmt.Sprintf("Error changing user's role in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Changed role for user with Username %s to %s in DB", username, newRole)
	return int(rowsAffected), nil
}

// ChangeUserPasswordByUserID updates a user's password in the MySQL database by user ID and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserPasswordByUserID(userID int, newPassword string) (int, error) {
	query := "UPDATE User SET Password = ? WHERE IDUser = ?"
	result, err := db.DB.Exec(query, newPassword, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Error changing user's password in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Changed password for user with ID %d in DB", userID)
	return int(rowsAffected), nil
}

// ChangeUserPasswordByUsername updates a user's password in the MySQL database by username and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserPasswordByUsername(username string, newPassword string) (int, error) {
	query := "UPDATE User SET Password = ? WHERE Username = ?"
	result, err := db.DB.Exec(query, newPassword, username)
	if err != nil {
		errMsg := fmt.Sprintf("Error changing user's password in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Changed password for user with Username %s in DB", username)
	return int(rowsAffected), nil
}

// UpdateUserTokenByID updates a user's authentication token by their ID and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserTokenByID(userID int, newToken string) (int, error) {
	query := "UPDATE User SET Token = ? WHERE IDUser = ?"
	result, err := db.DB.Exec(query, newToken, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Error updating user token in DB: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		errMsg := fmt.Sprintf("Error getting rows affected: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	log.Printf("[IDM] Updated token for user with ID %d in DB", userID)
	return int(rowsAffected), nil
}

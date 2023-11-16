package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// GetUserPasswordByUsername retrieves the hashed password of a user from the MySQL database by username.
func (db *MySQLDatabase) GetUserPasswordByUsername(username string) (string, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", utils.ColumnUserPassword, utils.UserTable, utils.ColumnUserName)

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
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", utils.ColumnRole, utils.RoleTable, utils.ColumnRoleIDUser)

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
	// query := "SELECT Role FROM Role r JOIN User u ON r.IDUser = u.IDUser WHERE u.Username = ?"
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

// ChangeUserRoleByUserID updates a user's role in the MySQL database by user ID and returns the number of rows affected.
func (db *MySQLDatabase) UpdateUserRoleByUserID(userID int, newRole string) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.RoleTable,
		utils.ColumnRole,
		utils.ColumnRoleIDUser,
	)
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
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = (SELECT %s FROM %s WHERE %s = ?)",
		utils.RoleTable,
		utils.ColumnRole,
		utils.ColumnRoleIDUser,
		utils.ColumnIDUser,
		utils.UserTable,
		utils.ColumnUserName,
	)
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
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserPassword,
		utils.ColumnIDUser,
	)
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
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?",
		utils.UserTable,
		utils.ColumnUserPassword,
		utils.ColumnUserName,
	)
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

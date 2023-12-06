package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

func (db *MySQLDatabase) AddUserToTrash(ctx context.Context, userData models.TrashData) error {
	// Construct the SQL insert query
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) VALUES (?, ?, ?, ?)", utils.TrashTable, utils.ColumnIDUser, utils.ColumnUserName, utils.ColumnUserPassword, utils.ColumnRole)

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	log.Println("[IDM] Attempting to insert user in trash")

	// Execute the SQL statement
	result, err := stmt.ExecContext(ctx, userData.IDUser, userData.Username, userData.Password, userData.Role)
	if err != nil {
		log.Printf("[IDM] Error executing query to insert user: %v", err)
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[IDM] Error getting last insert id when inserting user into trash: %v", err)
		return err
	}

	if lastInsertID == 0 {
		log.Printf("[IDM] Something unexpected happened and the user could not be moved to the trash.")
	} else {
		log.Printf("[IDM] User data moved to trash successfully. ID: %d", lastInsertID)
	}

	return nil
}

func (db *MySQLDatabase) GetDataFromTrashByUserID(ctx context.Context, userID int) (*models.TrashData, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ?",
		utils.ColumnIDUser,
		utils.ColumnUserName,
		utils.ColumnUserPassword,
		utils.ColumnRole,
		utils.TrashTable,
		utils.ColumnIDUser,
	)

	log.Printf("[IDM] Attempting to fetch user with ID %d from trash", userID)

	// Execute the SQL query with context
	row := db.QueryRowContext(ctx, query, userID)

	var userData models.TrashData
	err := row.Scan(&userData.IDUser, &userData.Username, &userData.Password, &userData.Role)
	if err != nil {
		log.Printf("[IDM] Error fetching user by ID %d from trash: %v", userID, err)
		return nil, err
	}

	log.Printf("[IDM] Successfully fetched user by ID %d.", userID)
	return &userData, nil
}

func (db *MySQLDatabase) RemoveUserFromTrash(ctx context.Context, userID int) (int, error) {
	// Construct the SQL insert query
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?",
		utils.TrashTable,
		utils.ColumnIDUser,
	)

	log.Printf("[IDM] Attempting to remove user with ID %d from trash", userID)

	// Execute the SQL query to delete the user
	result, err := db.DB.ExecContext(ctx, query, userID)
	if err != nil {
		// Log and return an error if the deletion fails
		errMsg := fmt.Sprintf("Error removing user from Trash: %v", err)
		log.Printf("[IDM] %s", errMsg)
		return 0, err
	}

	// Get the number of rows affected by the deletion
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[IDM] Error getting rows affected in Trash Table: %v", err)
		return 0, fmt.Errorf("error getting rows affected in Trash Table: %v", err)
	}

	// Check if no rows were affected, indicating that the user was not found
	if rowsAffected == 0 {
		log.Printf("[IDM] User with ID %d not found in trash", userID)
		return 0, nil
	}

	// Log that the user has been deleted successfully
	log.Printf("[IDM] Removed user from Trash: %d rows affected", rowsAffected)
	return int(rowsAffected), nil
}

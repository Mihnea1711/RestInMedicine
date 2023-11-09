package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// GetUsers retrieves all users.
func (c *IDMController) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Call the database method to retrieve all users
	users, err := c.DbConn.GetAllUsers()
	if err != nil {
		log.Printf("[IDM] Error getting all users: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	// Return the list of users in the response
	utils.RespondWithJSON(w, http.StatusOK, users)
}

// GetUserByID retrieves a user by ID.
func (c *IDMController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL parameters
	userID := mux.Vars(r)["id"]

	// Convert the userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("[IDM] Invalid user ID: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Call the database method to retrieve the user by ID
	user, err := c.DbConn.GetUserByID(userIDInt)
	if err != nil {
		log.Printf("[IDM] Error getting user by ID: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	if user.IDUser == 0 {
		// User not found
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	// Return the user in the response
	utils.RespondWithJSON(w, http.StatusOK, user)
}

// UpdateUserByID updates a user by ID.
func (c *IDMController) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL parameters
	userID := mux.Vars(r)["id"]

	// Convert the userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("[IDM] Invalid user ID: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Parse the updated user data from the request body
	var userCredentials models.CredentialsRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&userCredentials)
	if err != nil {
		log.Printf("[IDM] Error decoding request body: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Hash the updated password if it's present
	if userCredentials.Password != "" {
		hashedPassword, err := utils.HashPassword(userCredentials.Password)
		if err != nil {
			log.Printf("[IDM] Error hashing password: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}
		userCredentials.Password = hashedPassword
	}

	// Call the database method to update the user by ID
	rowsAffected, err := c.DbConn.UpdateUserByID(userCredentials, userIDInt)
	if err != nil {
		log.Printf("[IDM] Error updating user by ID: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found or no changes made"})
		return
	}

	// Return a success response
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User updated successfully", "rows_affected": fmt.Sprint(rowsAffected)})
}

// DeleteUserByID deletes a user by ID.
func (c *IDMController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL parameters
	userID := mux.Vars(r)["id"]

	// Convert the userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("[IDM] Invalid user ID: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Call the database method to delete the user by ID
	rowsAffected, err := c.DbConn.DeleteUserByID(userIDInt)
	if err != nil {
		log.Printf("[IDM] Error deleting user by ID: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	// Return a success response
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully", "rows_affected": fmt.Sprint(rowsAffected)})
}

// GetUserRole retrieves the role of a user.
func (c *IDMController) GetUserRole(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL parameters
	userID := mux.Vars(r)["id"]

	// Convert the userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("[IDM] Invalid user ID: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Call the database method to retrieve the user's role by ID
	userRole, err := c.DbConn.GetUserRoleByUserID(userIDInt)
	if err != nil {
		log.Printf("[IDM] Error getting user's role: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	if userRole == "" {
		// Role not found
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found or role not assigned"})
		return
	}

	// Return the user's role in the response
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"role": userRole})
}

// UpdateUserRole changes a user's role.
func (c *IDMController) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL parameters
	userID := mux.Vars(r)["id"]

	// Convert the userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("[IDM] Invalid user ID: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Parse the new role from the request body
	var newRole struct {
		Role string `json:"role"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&newRole)
	if err != nil {
		log.Printf("[IDM] Error decoding request body: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Call the database method to change the user's role
	rowsAffected, err := c.DbConn.UpdateUserRoleByUserID(userIDInt, newRole.Role)
	if err != nil {
		log.Printf("[IDM] Error changing user's role: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}
	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	// Return a success response
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Role changed successfully"})
}

// UpdateUserPassword changes a user's password.
func (c *IDMController) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL parameters
	userID := mux.Vars(r)["id"]

	// Convert the userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("[IDM] Invalid user ID: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Parse the new password from the request body
	var newPassword struct {
		Password string `json:"password"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&newPassword)
	if err != nil {
		log.Printf("[IDM] Error decoding request body: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Hash the new password before updating it
	hashedPassword, err := utils.HashPassword(newPassword.Password)
	if err != nil {
		log.Printf("[IDM] Error hashing password: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	// Call the database method to change the user's password
	rowsAffected, err := c.DbConn.UpdateUserPasswordByUserID(userIDInt, hashedPassword)
	if err != nil {
		log.Printf("[IDM] Error changing user's password: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}
	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	// Return a success response
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Password changed successfully"})
}

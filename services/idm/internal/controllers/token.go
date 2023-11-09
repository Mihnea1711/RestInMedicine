package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// UpdateUserToken updates a user's authentication token.
func (c *IDMController) UpdateUserToken(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL parameters
	userID := mux.Vars(r)["id"]

	// Convert the userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("[IDM] Invalid user ID: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Parse the new token from the request body
	var newToken struct {
		Token string `json:"token"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&newToken)
	if err != nil {
		log.Printf("[IDM] Error decoding request body: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Call the database method to update the user's token and get the rows affected
	rowsAffected, err := c.DbConn.UpdateUserTokenByID(userIDInt, newToken.Token)
	if err != nil {
		log.Printf("[IDM] Error updating user token: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	// Check if any rows were affected
	if rowsAffected == 0 {
		log.Printf("[IDM] User with ID %d not found or token unchanged", userIDInt)
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found or token unchanged"})
		return
	}

	// Return a success response
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Token updated successfully"})
}

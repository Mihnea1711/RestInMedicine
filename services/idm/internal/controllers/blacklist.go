package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// AddUserToBlacklist adds a user to the blacklist.
func (c *IDMController) AddUserToBlacklist(w http.ResponseWriter, r *http.Request) {
	var blacklistUserModel models.BlacklistToken
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&blacklistUserModel)
	if err != nil {
		log.Printf("[IDM] Error decoding request body for blacklisting: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Use the Redis client to add the user to the blacklist.
	err = c.RedisConn.AddUserToBlacklistInRedis(r.Context(), blacklistUserModel)
	if err != nil {
		log.Printf("[IDM] Error adding user to blacklist: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to add user to blacklist"})
		return
	}

	// Handle a successful addition to the blacklist.
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User added to the blacklist successfully"})
}

// RemoveUserFromBlacklist removes a user from the blacklist and returns the number of rows affected.
func (c *IDMController) RemoveUserFromBlacklist(w http.ResponseWriter, r *http.Request) {
	var blacklistUserModel models.BlacklistToken
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&blacklistUserModel)
	if err != nil {
		log.Printf("[IDM] Error decoding request body for blacklisting: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Use the Redis client to remove the user from the blacklist and get the number of rows affected.
	rowsAffected, err := c.RedisConn.RemoveUserFromBlacklistInRedis(r.Context(), blacklistUserModel)
	if err != nil {
		log.Printf("[IDM] Error removing user from blacklist: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to remove user from blacklist"})
		return
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not in the blacklist.
		log.Printf("[IDM] User with ID %d is not in the blacklist", blacklistUserModel.IDUser)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "User not found in the blacklist"})
		return
	}

	// Handle a successful removal from the blacklist.
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "User removed from the blacklist successfully", "rows_affected": rowsAffected})
}

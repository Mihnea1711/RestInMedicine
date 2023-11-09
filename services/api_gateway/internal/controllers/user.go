package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// RegisterUser handles user registration.
func (gc *GatewayController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Extract the role query parameter from the request
	role := r.URL.Query().Get("role")

	utils.RespondWithJSON(w, http.StatusOK, role)
}

func (c *GatewayController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUserRequest models.LoginUserRequest

	// Parse the request body into the LoginUserRequest struct
	if err := json.NewDecoder(r.Body).Decode(&loginUserRequest); err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, loginUserRequest)
}

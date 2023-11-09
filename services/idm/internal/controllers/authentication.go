package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// RegisterUser registers a new user.
func (c *IDMController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request to obtain user registration data
	user, ok := r.Context().Value(utils.DECODED_IDM).(models.UserRegistration)
	if !ok {
		log.Println("[IDM] Error retrieving user info from context")
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid user information"})
		return
	}

	// Hash the user's password before adding to the database
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("[IDM] Error hashing password: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	// Set the hashed password back to the user registration
	user.Password = hashedPassword

	// Call the database method to add the user to the database
	lastUserID, err := c.DbConn.AddUser(user)
	if err != nil {
		log.Printf("[IDM] Error adding user to the database: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	if lastUserID == 0 {
		// No rows were affected, which means the user was not added
		utils.RespondWithJSON(w, http.StatusConflict, map[string]string{"error": "User not added"})
		return
	}

	// Return a success response with the JWT token
	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "User registered successfully. Proceed to login.",
	})
}

// LoginUser handles user login.
func (c *IDMController) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request to obtain user login data
	userCredentials, ok := r.Context().Value(utils.DECODED_IDM).(models.CredentialsRequest)
	if !ok {
		log.Println("[IDM] Error retrieving user info from context")
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid user information"})
		return
	}

	// Retrieve the hashed password for the user from the database
	hashedPassword, err := c.DbConn.GetUserPasswordByUsername(userCredentials.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in the database: %s", userCredentials.Username)
			utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		} else {
			log.Printf("[IDM] Error retrieving user's hashed password: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		return
	}

	// Verify the password
	err = utils.VerifyPassword(hashedPassword, userCredentials.Password)
	if err != nil {
		log.Printf("[IDM] Invalid password for user: %s", userCredentials.Username)
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		return
	}

	// Retrieve the user's role
	userRole, err := c.DbConn.GetUserRoleByUsername(userCredentials.Username)
	if err != nil {
		log.Printf("[IDM] Error retrieving user's role: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	userComplete, err := c.DbConn.GetUserByUsername(userCredentials.Username)
	if err != nil {
		log.Printf("[IDM] Error retrieving user info: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	// Generate a JWT token
	token, err := utils.CreateJWT(userComplete.IDUser, userRole, c.jwtconfig)
	if err != nil {
		log.Printf("[IDM] Error generating JWT: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	// Return a success response with the JWT token
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}

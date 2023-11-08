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

	// You may want to validate user registration data here.

	// Hash the user's password before adding to the database
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("[IDM] Error hashing password: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	// Set the hashed password back to the user registration
	user.Password = hashedPassword

	// Generate a JWT token
	token, err := utils.CreateJWT(3, user.Role, c.jwtconfig)
	if err != nil {
		log.Printf("[IDM] Error generating JWT: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	// Call the database method to add the user to the database
	rowsAffected, err := c.DbConn.AddUser(user, token)
	if err != nil {
		log.Printf("[IDM] Error adding user to the database: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not added
		utils.RespondWithJSON(w, http.StatusConflict, map[string]string{"error": "User not added"})
		return
	}

	// Return a success response with the JWT token
	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "User registered successfully",
		"token":   token,
	})
}

// LoginUser handles user login.
func (c *IDMController) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request to obtain user login data
	user, ok := r.Context().Value(utils.DECODED_IDM).(models.User)
	if !ok {
		log.Println("[IDM] Error retrieving user info from context")
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid user information"})
		return
	}

	// Retrieve the hashed password for the user from the database
	hashedPassword, err := c.DbConn.GetUserPasswordByUsername(user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in the database: %s", user.Username)
			utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		} else {
			log.Printf("[IDM] Error retrieving user's hashed password: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		return
	}

	// Verify the password
	err = utils.VerifyPassword(hashedPassword, user.Password)
	if err != nil {
		log.Printf("[IDM] Invalid password for user: %s", user.Username)
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		return
	}

	// Retrieve the user's role
	var userRole string
	if user.IDUser > 0 {
		role, err := c.DbConn.GetUserRoleByUserID(user.IDUser)
		if err != nil {
			log.Printf("[IDM] Error retrieving user's role: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}
		userRole = role
	} else {
		role, err := c.DbConn.GetUserRoleByUsername(user.Username)
		if err != nil {
			log.Printf("[IDM] Error retrieving user's role: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			return
		}
		userRole = role
	}

	// Generate a JWT token
	token, err := utils.CreateJWT(user.IDUser, userRole, c.jwtconfig)
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

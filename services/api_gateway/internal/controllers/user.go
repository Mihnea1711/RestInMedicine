package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// GetAllUsers handles fetching all users.
func (gc *GatewayController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Implementation for fetching all users.
	log.Println("[GATEWAY] Handling GetAllUsers request...")

	response, err := gc.IDMClient.GetUsers(r.Context(), &proto_files.EmptyRequest{})
	if err != nil {
		log.Println("[GATEWAY] Error fetching all users:", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	log.Println("[GATEWAY] GetAllUsers request handled successfully.")
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetByIDUser handles fetching a user by ID.
func (gc *GatewayController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Implementation for fetching a user by ID.
	log.Println("[GATEWAY] Handling GetUserByID request...")

	userIDString := mux.Vars(r)["userID"]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Implement gRPC call to fetch a user by ID from IDM server.
	response, err := gc.IDMClient.GetUserByID(r.Context(), &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: userID}})
	if err != nil {
		log.Println("[GATEWAY] Error fetching user by ID:", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	log.Println("[GATEWAY] GetUserByID request handled successfully.")
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// UpdateUser handles updating a user.
func (gc *GatewayController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a user.
	log.Println("[GATEWAY] Handling UpdateUser request...")

	userIDString := mux.Vars(r)["userID"]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Implement gRPC call to update a user in the IDM server.
	var userData models.UserData
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		log.Println("[GATEWAY] Invalid request:", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	response, err := gc.IDMClient.UpdateUserByID(r.Context(), &proto_files.UpdateUserRequest{
		UserData: &proto_files.UserData{
			UserID:   &proto_files.UserID{ID: userID},
			Username: userData.Username,
			// Other fields...
		},
	})
	if err != nil {
		log.Println("[GATEWAY] Error updating user:", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	log.Println("[GATEWAY] UpdateUser request handled successfully.")
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// DeleteUser handles deleting a user.
func (gc *GatewayController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a user.
	log.Println("[GATEWAY] Handling DeleteUser request...")

	userIDString := mux.Vars(r)["userID"]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Implement gRPC call to delete a user in the IDM server.
	response, err := gc.IDMClient.DeleteUserByID(r.Context(), &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: userID}})
	if err != nil {
		log.Println("[GATEWAY] Error deleting user:", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	log.Println("[GATEWAY] DeleteUser request handled successfully.")
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// UpdatePassword handles updating a user's password.
func (gc *GatewayController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a user's password.
	log.Println("[GATEWAY] Handling UpdatePassword request...")

	userIDString := mux.Vars(r)["userID"]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Implement gRPC call to update a user's password in the IDM server.
	var passwordData models.PasswordData
	if err := json.NewDecoder(r.Body).Decode(&passwordData); err != nil {
		log.Println("[GATEWAY] Invalid request:", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	response, err := gc.IDMClient.UpdateUserPassword(r.Context(), &proto_files.UpdatePasswordRequest{
		UserID:   &proto_files.UserID{ID: userID},
		Password: passwordData.Password,
	})
	if err != nil {
		log.Println("[GATEWAY] Error updating user password:", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	log.Println("[GATEWAY] UpdatePassword request handled successfully.")
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// UpdateRole handles updating a user's role.
func (gc *GatewayController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a user's role.
	log.Println("[GATEWAY] Handling UpdateRole request...")

	userIDString := mux.Vars(r)["userID"]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Extract role from the request body
	var roleData models.RoleData
	if err := json.NewDecoder(r.Body).Decode(&roleData); err != nil {
		log.Println("[GATEWAY] Invalid request:", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	// Call the gRPC service with the provided information
	response, err := gc.IDMClient.UpdateUserRole(ctx, &proto_files.UpdateRoleRequest{
		UserID: &proto_files.UserID{ID: userID},
		Role:   roleData.Role,
	})

	if err != nil {
		log.Println("[GATEWAY] Error updating user role:", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	log.Println("[GATEWAY] UpdateRole request handled successfully.")
	utils.RespondWithJSON(w, http.StatusOK, response)
}

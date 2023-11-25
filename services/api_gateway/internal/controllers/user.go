package controllers

import (
	"context"
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
	log.Println("[GATEWAY] Handling GetAllUsers request...")

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	response, err := gc.IDMClient.GetUsers(ctx, &proto_files.EmptyRequest{})
	if err != nil {
		log.Println("[GATEWAY] Error fetching all users:", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] Get All Users response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Get All Users response is nil", "Received nil response while getting all users.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] Get Users response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response.Info while getting users.")
		return
	}

	if response.Users == nil {
		log.Println("[GATEWAY] Users object is nil")
		utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "Received nil response.Users while getting users.")
		return
	}

	// Check the gRPC response status and handle accordingly
	switch response.Info.Status {
	case http.StatusOK:
		var users []models.UserData
		// Convert proto users to users
		for _, protoUser := range response.Users {
			user := models.UserData{
				IDUser:   int(protoUser.UserID.ID),
				Username: protoUser.Username,
			}
			users = append(users, user)
		}

		log.Println("[GATEWAY] GetAllUsers request handled successfully.")
		utils.SendMessageResponse(w, http.StatusOK, response.Info.Message, users)
		return
	default:
		// Other status codes
		log.Printf("[GATEWAY] Unexpected status code: %d", response.Info.Status)
		utils.SendMessageResponse(w, http.StatusInternalServerError, response.Info.Message, "An unexpected error occurred while getting all users. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}
}

// GetUserByID handles fetching a user by ID.
func (gc *GatewayController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("[GATEWAY] Handling GetUserByID request...")

	userIDString := mux.Vars(r)[utils.GET_USER_ID_PARAMETER]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Implement gRPC call to fetch a user by ID from IDM server.
	response, err := gc.IDMClient.GetUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: userID}})
	if err != nil {
		log.Printf("[GATEWAY] Error fetching user by ID: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] Get User By ID response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Get User By ID response is nil", "Received nil response while getting the user.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] Get User By ID response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response.Info while getting user by id.")
		return
	}

	if response.User == nil {
		log.Println("[GATEWAY] User not found")
		utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "Received nil response.User while getting the user.")
		return
	}

	// Check the gRPC response status and handle accordingly
	switch response.Info.Status {
	case http.StatusOK:
		user := models.UserData{
			IDUser:   int(response.User.UserID.ID),
			Username: response.User.Username,
			// Add other fields if needed
		}

		log.Println("[GATEWAY] GetUserByID request handled successfully.")
		utils.SendMessageResponse(w, http.StatusOK, response.Info.Message, user)
		return
	case http.StatusNotFound:
		// User not found
		log.Printf("[GATEWAY] User not found: %s", response.Info.Message)
		utils.SendMessageResponse(w, http.StatusNotFound, response.Info.Message, "The specified user was not found or no changes were made.")
		return
	default:
		// Other status codes
		log.Printf("[GATEWAY] Unexpected status code: %d", response.Info.Status)
		utils.SendMessageResponse(w, http.StatusInternalServerError, response.Info.Message, "An unexpected error occurred while updating user. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}
}

// UpdateUser handles updating a user.
func (gc *GatewayController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a user.
	log.Println("[GATEWAY] Handling UpdateUser request...")

	userIDString := mux.Vars(r)[utils.UPDATE_USER_ID_PARAMETER]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// var userData models.UserData
	// if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
	// 	log.Println("[GATEWAY] Invalid request:", err)
	// 	utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request", err.Error())
	// 	return
	// }
	// Take user data from the context after validation
	userData := r.Context().Value(utils.DECODED_USER_DATA).(*models.UserData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Implement gRPC call to update a user in the IDM server.
	response, err := gc.IDMClient.UpdateUserByID(ctx, &proto_files.UpdateUserRequest{
		UserData: &proto_files.UserData{
			UserID:   &proto_files.UserID{ID: userID},
			Username: userData.Username,
		},
	})
	if err != nil {
		log.Printf("[GATEWAY] Error updating user: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Failed to update user: "+err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] UpdateUser response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response while updating user.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] UpdateUser response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response.Info while updating user.")
		return
	}

	switch response.Info.Status {
	case http.StatusOK:
		log.Println("[GATEWAY] User updated successfully.")
		enhancedResponse := models.RowsAffected{
			RowsAffected: int(response.RowsAffected),
		}
		utils.SendMessageResponse(w, http.StatusOK, "User updated successfully.", enhancedResponse)
		return
	case http.StatusNotFound:
		log.Println("[GATEWAY] User not found or no changes made.")
		utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "The specified user was not found or no changes were made.")
		return
	default:
		log.Printf("[GATEWAY] UpdateUser failed with status %d: %s", response.Info.Status, response.Info.Message)
		utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, "An unexpected error occurred while updating user. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}
}

// DeleteUser handles deleting a user.
func (gc *GatewayController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a user.
	log.Println("[GATEWAY] Handling DeleteUser request...")

	userIDString := mux.Vars(r)[utils.DELETE_USER_ID_PARAMETER]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Implement gRPC call to delete a user in the IDM server.
	response, err := gc.IDMClient.DeleteUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: userID}})
	if err != nil {
		log.Printf("[GATEWAY] Error deleting user: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Failed to delete user: "+err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] DeleteUser response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response while deleting user.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] DeleteUser response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response.Info while deleting user.")
		return
	}

	switch response.Info.Status {
	case http.StatusOK:
		log.Println("[GATEWAY] User deleted successfully.")
		enhancedResponse := models.RowsAffected{
			RowsAffected: int(response.RowsAffected),
		}
		utils.SendMessageResponse(w, http.StatusOK, "User deleted successfully.", enhancedResponse)
		return
	case http.StatusNotFound:
		log.Println("[GATEWAY] User not found or no changes made.")
		utils.SendErrorResponse(w, http.StatusNotFound, "User not found or no changes made.", "")
		return
	default:
		log.Printf("[GATEWAY] DeleteUser failed with status %d: %s", response.Info.Status, response.Info.Message)
		utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, "DeleteUser failed. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}
}

// UpdatePassword handles updating a user's password.
func (gc *GatewayController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a user's password.
	log.Println("[GATEWAY] Handling UpdatePassword request...")

	userIDString := mux.Vars(r)[utils.UPDATE_USER_PASSWORD_ID_PARAMETER]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// // Implement gRPC call to update a user's password in the IDM server.
	// var passwordData models.PasswordData
	// if err := json.NewDecoder(r.Body).Decode(&passwordData); err != nil {
	// 	log.Println("[GATEWAY] Invalid request payload for changing password:", err)
	// 	utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload for changing password", err.Error())
	// 	return
	// }
	// Take password data from the context after validation
	passwordData := r.Context().Value(utils.DECODED_PASSWORD_DATA).(*models.PasswordData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	response, err := gc.IDMClient.UpdateUserPassword(ctx, &proto_files.UpdatePasswordRequest{
		UserID:   &proto_files.UserID{ID: userID},
		Password: passwordData.Password,
	})
	if err != nil {
		log.Printf("[GATEWAY] Error updating user password: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Failed to update user password: "+err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] UpdateUserPassword response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response while updating user password.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] UpdateUserPassword response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response.Info while updating user password.")
		return
	}

	switch response.Info.Status {
	case http.StatusOK:
		log.Println("[GATEWAY] User password updated successfully.")
		enhancedResponse := models.RowsAffected{
			RowsAffected: int(response.RowsAffected),
		}
		utils.SendMessageResponse(w, http.StatusOK, "User password updated successfully.", enhancedResponse)
		return
	case http.StatusNotFound:
		log.Println("[GATEWAY] User not found or no changes made.")
		utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "The user was not found or no changes were made to the password.")
		return
	default:
		log.Printf("[GATEWAY] UpdateUserPassword failed with status %d: %s", response.Info.Status, response.Info.Message)
		utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, "An unexpected error occurred while updating user password. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}
}

// UpdateRole handles updating a user's role.
func (gc *GatewayController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a user's role.
	log.Println("[GATEWAY] Handling UpdateRole request...")

	userIDString := mux.Vars(r)[utils.UPDATE_USER_ROLE_ID_PARAMETER]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// Extract role from the request body
	// var roleData models.RoleData
	// if err := json.NewDecoder(r.Body).Decode(&roleData); err != nil {
	// 	log.Println("[GATEWAY] Invalid request:", err)
	// 	utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload for changing role", err.Error())
	// 	return
	// }
	// Take role data from the context after validation
	roleData := r.Context().Value(utils.DECODED_ROLE_DATA).(*models.RoleData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Call the gRPC service with the provided information
	response, err := gc.IDMClient.UpdateUserRole(ctx, &proto_files.UpdateRoleRequest{
		UserID: &proto_files.UserID{ID: userID},
		Role:   roleData.Role,
	})

	if err != nil {
		log.Printf("[GATEWAY] Error updating user role: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Failed to update user role: "+err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] UpdateUserRole response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response while updating user role.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] UpdateUserRole response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response.Info while updating user role.")
		return
	}

	switch response.Info.Status {
	case http.StatusOK:
		log.Println("[GATEWAY] User role updated successfully.")
		enhancedResponse := models.RowsAffected{
			RowsAffected: int(response.RowsAffected),
		}
		utils.SendMessageResponse(w, http.StatusOK, "User role updated successfully.", enhancedResponse)
		return
	case http.StatusNotFound:
		log.Println("[GATEWAY] User not found or no changes made.")
		utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "The specified user was not found or no changes were made.")
		return
	default:
		log.Printf("[GATEWAY] UpdateUserRole failed with status %d: %s", response.Info.Status, response.Info.Message)
		utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, "An unexpected error occurred while updating user role. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}
}

// AddToBlacklist handles adding a user to the blacklist.
func (gc *GatewayController) AddToBlacklist(w http.ResponseWriter, r *http.Request) {
	log.Println("[GATEWAY] Handling AddToBlacklist request...")

	// Take blacklist data from the context after validation
	blacklistRequest := r.Context().Value(utils.DECODED_BLACKLIST_DATA).(*models.BlacklistData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Implement gRPC call to add a user to the blacklist in IDM server.
	response, err := gc.IDMClient.AddUserToBlacklist(ctx, &proto_files.BlacklistRequest{
		UserID: &proto_files.UserID{ID: int64(blacklistRequest.IDUser)},
		Token:  blacklistRequest.Token,
	})
	if err != nil {
		log.Printf("[GATEWAY] Error adding user to blacklist: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Failed to add user to the blacklist: "+err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] AddToBlacklist response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response while adding user to the blacklist.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] AddToBlacklist response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response.Info while adding user to the blacklist.")
		return
	}

	switch response.Info.Status {
	case http.StatusOK:
		log.Println("[GATEWAY] AddToBlacklist request handled successfully.")
		utils.SendMessageResponse(w, http.StatusOK, "User added to the blacklist successfully.", nil)
		return
	case http.StatusNotFound:
		log.Println("[GATEWAY] User not found or no changes made.")
		utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "The specified user was not found or no changes were made.")
		return
	case http.StatusConflict:
		log.Println("[GATEWAY] User is already in the blacklist.")
		utils.SendErrorResponse(w, http.StatusConflict, response.Info.Message, "The specified user was is already in the blacklsit. No changes were made.")
		return
	default:
		log.Printf("[GATEWAY] AddUserToBlacklist failed with status %d: %s", response.Info.Status, response.Info.Message)
		utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, "An unexpected error occurred while adding user to the blacklist. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}

}

// CheckBlacklist handles checking if a user is in the blacklist.
func (gc *GatewayController) CheckBlacklist(w http.ResponseWriter, r *http.Request) {
	log.Println("[GATEWAY] Handling CheckBlacklist request...")

	userIDString := mux.Vars(r)[utils.BLACKLIST_USER_ID_PARAMETER]
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Implement gRPC call to check if a user is in the blacklist in IDM server.
	response, err := gc.IDMClient.CheckUserInBlacklist(ctx, &proto_files.UserIDRequest{
		UserID: &proto_files.UserID{ID: userID},
	})
	if err != nil {
		log.Printf("[GATEWAY] Error checking user in blacklist: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Failed to check if the user is in the blacklist: "+err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] CheckBlacklist response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response while checking if the user is in the blacklist.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] CheckBlacklist response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response.Info while checking if the user is in the blacklist.")
		return
	}

	switch response.Info.Status {
	case http.StatusOK:
		log.Println("[GATEWAY] CheckBlacklist request handled successfully.")
		utils.SendMessageResponse(w, http.StatusOK, "User is in the blacklist.", nil)
		return
	case http.StatusNotFound:
		log.Println("[GATEWAY] User not found or no changes made.")
		utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "The specified user was not found or no changes were made.")
		return
	default:
		log.Printf("[GATEWAY] CheckUserInBlacklist failed with status %d: %s", response.Info.Status, response.Info.Message)
		utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, "An unexpected error occurred while checking if the user is in the blacklist. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}

}

// RemoveFromBlacklist handles removing a user from the blacklist.
func (gc *GatewayController) RemoveFromBlacklist(w http.ResponseWriter, r *http.Request) {
	log.Println("[GATEWAY] Handling RemoveFromBlacklist request...")

	userIDString := mux.Vars(r)[utils.DELETE_USER_ID_PARAMETER]
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println("[GATEWAY] Invalid user ID:", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Implement gRPC call to remove a user from the blacklist in IDM server.
	response, err := gc.IDMClient.RemoveUserFromBlacklist(ctx, &proto_files.UserIDRequest{
		UserID: &proto_files.UserID{ID: userID},
	})
	if err != nil {
		log.Printf("[GATEWAY] Error removing user from blacklist: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error removing user from blacklist: "+err.Error())
		return
	}

	if response == nil {
		log.Println("[GATEWAY] RemoveFromBlacklist response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "RemoveFromBlacklist response is nil", "The RemoveFromBlacklist response is nil.")
		return
	}

	if response.Info == nil {
		log.Println("[GATEWAY] RemoveFromBlacklist response.Info is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "RemoveFromBlacklist response.Info is nil", "The RemoveFromBlacklist response.Info is nil.")
		return
	}

	switch response.Info.Status {
	case http.StatusOK:
		log.Println("[GATEWAY] User role updated successfully.")
		enhancedResponse := models.RowsAffected{
			RowsAffected: int(response.RowsAffected),
		}
		utils.SendMessageResponse(w, http.StatusOK, response.Info.Message, enhancedResponse)
		return
	case http.StatusNotFound:
		log.Println("[GATEWAY] User not found or no changes made.")
		utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "The requested user was not found, or no changes were made")
		return
	default:
		log.Printf("[GATEWAY] RemoveFromBlacklist failed with status %d: %s", response.Info.Status, response.Info.Message)
		utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, "RemoveFromBlacklist failed. Unexpected status code: "+strconv.Itoa(int(response.Info.Status)))
		return
	}
}

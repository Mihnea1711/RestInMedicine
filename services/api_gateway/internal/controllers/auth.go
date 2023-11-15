package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// RegisterUser handles user registration.
func (gc *GatewayController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into the RegisterRequest struct
	var registerRequest models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		// Handle the error (e.g., return a response with an error message)
		log.Printf("[GATEWAY] Error decoding request body: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	log.Printf("[GATEWAY] Sending gRPC request to IDM for user registration: %+v", registerRequest)

	// Call the gRPC service with the provided information
	response, err := gc.IDMClient.Register(ctx, &proto_files.RegisterRequest{
		Role: registerRequest.Role,
		UserCredentials: &proto_files.UserCredenetials{
			Username: registerRequest.Username,
			Password: registerRequest.Password,
		},
	})

	if err != nil {
		// Handle gRPC error (e.g., return a response with an error message)
		log.Printf("[GATEWAY] Error calling IDM gRPC service: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	log.Printf("[GATEWAY] Received gRPC response from IDM: %+v", response)

	// Handle the gRPC response (e.g., return a response with the gRPC response)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// LoginUser handles user login.
func (gc *GatewayController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUserRequest models.LoginRequest

	// Parse the request body into the LoginUserRequest struct
	if err := json.NewDecoder(r.Body).Decode(&loginUserRequest); err != nil {
		// Handle the error (e.g., return a response with an error message)
		log.Printf("[GATEWAY] Error decoding request body: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	log.Printf("[GATEWAY] Sending gRPC request to IDM for login: %+v", loginUserRequest)

	// Call the gRPC service with the provided information
	response, err := gc.IDMClient.Login(ctx, &proto_files.LoginRequest{
		UserCredentials: &proto_files.UserCredenetials{
			Username: loginUserRequest.Username,
			Password: loginUserRequest.Password,
		},
	})

	if err != nil {
		// Handle gRPC error (e.g., return a response with an error message)
		log.Printf("[GATEWAY] Error calling IDM gRPC service: %v", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	log.Printf("[GATEWAY] Received gRPC response from IDM: %+v", response)

	// Handle the gRPC response (e.g., return a response with the gRPC response)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

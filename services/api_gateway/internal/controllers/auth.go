package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// RegisterUser handles user registration.
func (gc *GatewayController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to register a new user.")
	registerRequest := r.Context().Value(utils.DECODED_USER_REGISTRATION_DATA).(*models.UserRegistrationData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

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
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
	switch response.Info.Status {
	case http.StatusOK:
		// Registration successful
		log.Println("[GATEWAY] Registration successful. Proceed to login.")
		utils.SendMessageResponse(w, http.StatusOK, response.Info.Message, nil)
	case http.StatusConflict:
		// User not added
		log.Printf("[GATEWAY] Registration failed: %s", response.Info.Message)
		utils.SendMessageResponse(w, http.StatusConflict, response.Info.Message, nil)
	default:
		// Other status codes
		log.Printf("[GATEWAY] Registration failed: %s", response.Info.Message)
		utils.SendMessageResponse(w, int(response.Info.Status), response.Info.Message, nil)
	}
}

// LoginUser handles user login.
func (gc *GatewayController) LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to login a user.")
	loginUserRequest := r.Context().Value(utils.DECODED_USER_LOGIN_DATA).(*models.UserLoginData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

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
		utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
	switch response.Info.Status {
	case http.StatusOK:
		// Login successful
		log.Println("[GATEWAY] Login successful.")

		// Set the JWT token in a cookie
		cookie := http.Cookie{
			Name:     utils.COOKIE_NAME,
			Value:    response.Token, // Assuming Token is the field in your response containing the JWT token
			Path:     utils.COOKIE_PATH,
			MaxAge:   utils.COOKIE_MAX_AGE, // Max age of the cookie in seconds (e.g., 1 hour)
			HttpOnly: true,                 // This is important for security; prevents JavaScript access to the cookie
		}
		http.SetCookie(w, &cookie)

		utils.SendMessageResponse(w, http.StatusOK, response.Info.Message, nil)
	case http.StatusNotFound:
		// User not found
		log.Printf("[GATEWAY] User not found: %s", response.Info.Message)
		utils.SendMessageResponse(w, http.StatusNotFound, response.Info.Message, nil)
	case http.StatusUnauthorized:
		// Invalid credentials
		log.Printf("[GATEWAY] Invalid credentials: %s", response.Info.Message)
		utils.SendMessageResponse(w, http.StatusUnauthorized, response.Info.Message, nil)
	default:
		// Other status codes
		log.Printf("[GATEWAY] Unexpected status code: %d", response.Info.Status)
		utils.SendMessageResponse(w, http.StatusUnauthorized, response.Info.Message, nil)
	}
}

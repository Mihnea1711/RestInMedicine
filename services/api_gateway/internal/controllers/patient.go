package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// CreatePacient handles the creation of a new pacient.
func (gc *GatewayController) CreatePacient(w http.ResponseWriter, r *http.Request) {
	var pacientRequest models.PacientData

	// Parse the request body into the PacientData struct
	if err := json.NewDecoder(r.Body).Decode(&pacientRequest); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// check IDUser exists
	userResponse, err := gc.IDMClient.GetUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(pacientRequest.IDUser)}})
	if err != nil {
		log.Printf("[GATEWAY] Error fetching user by ID: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	if userResponse == nil {
		log.Println("[GATEWAY] Get User By ID response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Get User By ID response is nil", "")
		return
	}

	if userResponse.User == nil {
		log.Println("[GATEWAY] User does not exist")
		utils.SendErrorResponse(w, http.StatusConflict, "User does not exist", "")
		return
	}

	//

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, utils.POST, utils.CREATE_PATIENT_ENDPOINT, utils.PATIENT_PORT, pacientRequest)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
		return
	}

	// Respond with the response from the other module
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "Patient created successfully", "patient_data": pacientRequest, "response": response})
}

// GetPacienti handles fetching all pacients.
func (gc *GatewayController) GetPacienti(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request to another module
	response, err := gc.redirectRequestBody(ctx, utils.GET, utils.GET_ALL_PATIENTS_ENDPOINT, utils.PATIENT_PORT, nil)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
		return
	}

	// Respond with the response from the other module
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "Pacienti fetched successfully", "response": response})
}

// GetPacientByID handles fetching a pacient by ID.
func (gc *GatewayController) GetPacientByID(w http.ResponseWriter, r *http.Request) {
	pacientIDString := mux.Vars(r)["id"]

	// Convert pacientIDString to int64
	pacientID, err := strconv.ParseInt(pacientIDString, 10, 64)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid pacient ID"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s%d", utils.GET_PATIENT_BY_ID_ENDPOINT, pacientID), utils.PATIENT_PORT, nil)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
		return
	}

	// Respond with the response from the other module
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetPacientByEmail handles fetching a pacient by email.
func (gc *GatewayController) GetPacientByEmail(w http.ResponseWriter, r *http.Request) {
	pacientEmail := mux.Vars(r)["email"]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s%s", utils.GET_PATIENT_BY_EMAIL_ENDPOINT, pacientEmail), utils.PATIENT_PORT, nil)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
		return
	}

	// Respond with the response from the other module
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetPacientByUserID handles fetching a pacient by user ID.
// func (gc *GatewayController) GetPacientByUserID(w http.ResponseWriter, r *http.Request) {
// 	userIDString := mux.Vars(r)["id"]

// 	// Convert userIDString to int64
// 	userID, err := strconv.ParseInt(userIDString, 10, 64)
// 	if err != nil {
// 		// Handle the error (e.g., return a response with an error message)
// 		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
// 		return
// 	}

// 	// Redirect the request body to another module
// 	response, err := gc.redirectRequestBody(http.MethodGet, fmt.Sprintf("%s%d", utils.GET_PATIENT_BY_USER_ID_ENDPOINT, userID), utils.PATIENT_PORT, nil)
// 	if err != nil {
// 		// Handle the error (e.g., return a response with an error message)
// 		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
// 		return
// 	}

// 	// Respond with the response from the other module
// 	utils.RespondWithJSON(w, http.StatusOK, response)
// }

// UpdatePacientByID handles updating a pacient by ID.
func (gc *GatewayController) UpdatePacientByID(w http.ResponseWriter, r *http.Request) {
	pacientIDString := mux.Vars(r)["id"]

	// Convert pacientIDString to int64
	pacientID, err := strconv.ParseInt(pacientIDString, 10, 64)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid pacient ID"})
		return
	}

	var pacientData models.PacientData
	if err := json.NewDecoder(r.Body).Decode(&pacientData); err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodPut, fmt.Sprintf("%s%d", utils.UPDATE_PATIENT_BY_ID_ENDPOINT, pacientID), utils.PATIENT_PORT, pacientData)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
		return
	}

	// Respond with the response from the other module
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// DeletePacientByID handles deleting a pacient by ID.
func (gc *GatewayController) DeletePacientByID(w http.ResponseWriter, r *http.Request) {
	pacientIDString := mux.Vars(r)["id"]

	// Convert pacientIDString to int64
	pacientID, err := strconv.ParseInt(pacientIDString, 10, 64)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid pacient ID"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodDelete, fmt.Sprintf("%s%d", utils.DELETE_PATIENT_BY_ID_ENDPOINT, pacientID), utils.PATIENT_PORT, nil)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
		return
	}

	// Respond with the response from the other module
	utils.RespondWithJSON(w, http.StatusOK, response)
}

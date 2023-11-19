package controllers

import (
	"context"
	"encoding/json"
	"errors"
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

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, utils.POST, utils.PATIENT_CREATE_PATIENT_ENDPOINT, utils.PATIENT_PORT, pacientRequest)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			var responseBody *models.ResponseData
			if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
				log.Printf("[GATEWAY] Error decoding response body: %v", err)
				utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
				return
			}
			// Respond with the response from the other module
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusConflict:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, "Conflict in the request.", "")
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Something unexpected happened")
		return
	}
}

// GetPacienti handles fetching all pacients.
func (gc *GatewayController) GetPacienti(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request to another module
	response, err := gc.redirectRequestBody(ctx, utils.GET, utils.PATIENT_FETCH_ALL_PATIENTS_ENDPOINT, utils.PATIENT_PORT, nil)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		}
	}()

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Patients fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// GetPacientByID handles fetching a pacient by ID.
func (gc *GatewayController) GetPacientByID(w http.ResponseWriter, r *http.Request) {
	pacientIDString := mux.Vars(r)[utils.GET_PATIENT_ID_PARAMETER]

	// Convert pacientIDString to int64
	pacientID, err := strconv.ParseInt(pacientIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid pacient ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, pacientID), utils.PATIENT_PORT, nil)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		}
	}()

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Patient fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, "Patient Not Found.", responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// GetPacientByEmail handles fetching a pacient by email.
func (gc *GatewayController) GetPacientByEmail(w http.ResponseWriter, r *http.Request) {
	pacientEmail := mux.Vars(r)[utils.GET_PATIENT_EMAIL_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.PATIENT_FETCH_PATIENT_BY_EMAIL_ENDPOINT, pacientEmail), utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		}
	}()

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Patient fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, "Patient Not Found.", responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// GetPacientByUserID handles fetching a pacient by user ID.
func (gc *GatewayController) GetPacientByUserID(w http.ResponseWriter, r *http.Request) {
	userIDString := mux.Vars(r)[utils.GET_PATIENT_USER_ID_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid user ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_USER_ID_ENDPOINT, userID), utils.PATIENT_PORT, nil)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}
	}()

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Patient fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, "Patient Not Found.", responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// UpdatePacientByID handles updating a pacient by ID.
func (gc *GatewayController) UpdatePacientByID(w http.ResponseWriter, r *http.Request) {
	pacientIDString := mux.Vars(r)[utils.UPDATE_PATIENT_ID_PARAMETER]

	// Convert pacientIDString to int64
	pacientID, err := strconv.ParseInt(pacientIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Error closing response body: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid pacient ID", err.Error())
		return
	}

	var pacientData models.PacientData
	if err := json.NewDecoder(r.Body).Decode(&pacientData); err != nil {
		log.Printf("[GATEWAY] Error closing response body: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodPut, fmt.Sprintf("%s/%d", utils.PATIENT_UPDATE_PATIENT_BY_ID_ENDPOINT, pacientID), utils.PATIENT_PORT, pacientData)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}
	}()

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Patient updated successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle not found case
			utils.SendErrorResponse(w, http.StatusConflict, "Patient Not Found.", responseBody.Error)
			return
		}
	case http.StatusConflict:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, "Patient Data Conflict.", responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// DeletePacientByID handles deleting a pacient by ID.
func (gc *GatewayController) DeletePacientByID(w http.ResponseWriter, r *http.Request) {
	pacientIDString := mux.Vars(r)[utils.DELETE_PATIENT_ID_PARAMETER]

	// Convert pacientIDString to int64
	pacientID, err := strconv.ParseInt(pacientIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid pacient ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid pacient ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodDelete, fmt.Sprintf("%s/%d", utils.PATIENT_DELETE_PATIENT_BY_ID_ENDPOINT, pacientID), utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}
	}()

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Patient deleted successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle not found case
			utils.SendErrorResponse(w, http.StatusConflict, "Patient Not Found.", responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

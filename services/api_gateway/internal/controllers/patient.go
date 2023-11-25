package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// CreatePatient handles the creation of a new patient.
func (gc *GatewayController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to create a patient.")
	// Take credentials data from the context after validation
	patientRequest := r.Context().Value(utils.DECODED_PATIENT_DATA).(*models.PatientData)

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if IDUser exists
	userResponse, err := gc.IDMClient.GetUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(patientRequest.IDUser)}})
	if err != nil {
		log.Printf("[GATEWAY] Error fetching user by ID: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Failed to fetch user by ID: "+err.Error())
		return
	}

	if userResponse == nil {
		log.Println("[GATEWAY] Get User By ID response is nil")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Received nil response while fetching user by ID.")
		return
	}

	if userResponse.User == nil {
		log.Println("[GATEWAY] User does not exist")
		utils.SendErrorResponse(w, http.StatusNotFound, "User does not exist", "The specified user does not exist.")
		return
	}

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, utils.POST, utils.PATIENT_CREATE_PATIENT_ENDPOINT, utils.PATIENT_PORT, patientRequest)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", cerr)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error closing response body: "+cerr.Error())
		}
	}()

	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[GATEWAY] Error reading response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
		return
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseData
	if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		// Respond with the response from the other module
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusConflict:
		// Handle conflict case
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, decodedResponse.Error)
		return
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, decodedResponse.Error)
		return
	}
}

// GetPatients handles fetching all patients.
func (gc *GatewayController) GetPatients(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request to another module
	response, err := gc.redirectRequestBody(ctx, utils.GET, utils.PATIENT_FETCH_ALL_PATIENTS_ENDPOINT, utils.PATIENT_PORT, nil)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		log.Printf("[GATEWAY] Error redirecting request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", cerr)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error closing response body: "+cerr.Error())
		}
	}()

	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[GATEWAY] Error reading response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
		return
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseData
	if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		log.Println("[GATEWAY] Patients fetched successfully")
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetPatientByID handles fetching a patient by ID.
func (gc *GatewayController) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	patientIDString := mux.Vars(r)[utils.GET_PATIENT_ID_PARAMETER]

	// Convert patientIDString to int64
	patientID, err := strconv.ParseInt(patientIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid patient ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, patientID), utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", cerr)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error closing response body: "+cerr.Error())
		}
	}()

	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[GATEWAY] Error reading response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
		return
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseData
	if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		log.Println("[GATEWAY] Patient fetched successfully")
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		// Handle not found case
		log.Println("[GATEWAY] Patient not found")
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
		return
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetPatientByEmail handles fetching a patient by email.
func (gc *GatewayController) GetPatientByEmail(w http.ResponseWriter, r *http.Request) {
	patientEmail := mux.Vars(r)[utils.GET_PATIENT_EMAIL_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.PATIENT_FETCH_PATIENT_BY_EMAIL_ENDPOINT, patientEmail), utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}

	// Close the response body explicitly after decoding
	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", cerr)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error closing response body: "+cerr.Error())
		}
	}()

	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[GATEWAY] Error reading response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
		return
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseData
	if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		log.Println("[GATEWAY] Patient fetched successfully")
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		// Handle not found case
		log.Println("[GATEWAY] Patient not found")
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
		return
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetPatientByUserID handles fetching a patient by user ID.
func (gc *GatewayController) GetPatientByUserID(w http.ResponseWriter, r *http.Request) {
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
		}
	}()

	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[GATEWAY] Error reading response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
		return
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseData
	if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		log.Println("[GATEWAY] Patient fetched successfully")
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		// Handle conflict case
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
		return
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// UpdatePatientByID handles updating a patient by ID.
func (gc *GatewayController) UpdatePatientByID(w http.ResponseWriter, r *http.Request) {
	patientIDString := mux.Vars(r)[utils.UPDATE_PATIENT_ID_PARAMETER]

	// Convert patientIDString to int64
	patientID, err := strconv.ParseInt(patientIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Error closing response body: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid patient ID", err.Error())
		return
	}

	// Take patient data from context
	patientData := r.Context().Value(utils.DECODED_PATIENT_DATA).(*models.PatientData)

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodPut, fmt.Sprintf("%s/%d", utils.PATIENT_UPDATE_PATIENT_BY_ID_ENDPOINT, patientID), utils.PATIENT_PORT, patientData)
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

	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[GATEWAY] Error reading response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
		return
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseData
	if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Patient updated successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle not found case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
			return
		}
	case http.StatusConflict:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Patient Conflict: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// DeletePatientByID handles deleting a patient by ID.
func (gc *GatewayController) DeletePatientByID(w http.ResponseWriter, r *http.Request) {
	patientIDString := mux.Vars(r)[utils.DELETE_PATIENT_ID_PARAMETER]

	// Convert patientIDString to int64
	patientID, err := strconv.ParseInt(patientIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid patient ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid patient ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodDelete, fmt.Sprintf("%s/%d", utils.PATIENT_DELETE_PATIENT_BY_ID_ENDPOINT, patientID), utils.PATIENT_PORT, nil)
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

	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[GATEWAY] Error reading response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
		return
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseData
	if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Patient deleted successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle not found case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

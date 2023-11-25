package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils/wrappers"
)

// CreatePatient handles the creation of a new patient.
func (gc *GatewayController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to create a patient.")

	// Take credentials data from the context after validation
	patientRequest := r.Context().Value(utils.DECODED_PATIENT_DATA).(*models.PatientData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if IDUser exists
	userResponse, err := gc.IDMClient.GetUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(patientRequest.IDUser)}})
	if err != nil {
		log.Printf("[GATEWAY] Error fetching user by ID: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Failed to fetch user by ID: "+err.Error())
		return
	}
	// Check response for nils
	userResponseWrapper := &wrappers.UserResponse{Response: userResponse}
	utils.CheckNilResponse(w, http.StatusInternalServerError, "Get User By ID response is nil", userResponseWrapper.IsResponseNil, "Received nil response while getting the user.")
	utils.CheckNilResponse(w, http.StatusInternalServerError, "Get User By ID response info is nil", userResponseWrapper.IsInfoNil, "Received nil response.Info while getting user by id.")
	utils.CheckNilResponse(w, http.StatusInternalServerError, userResponse.Info.Message, userResponseWrapper.IsUserNil, "Received nil response.User while getting the user.")

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.POST, utils.PATIENT_CREATE_PATIENT_ENDPOINT, utils.PATIENT_PORT, patientRequest)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}
	// // Close the response body explicitly after decoding
	// defer func() {
	// 	if cerr := response.Body.Close(); cerr != nil {
	// 		log.Printf("[GATEWAY] Error closing response body: %v", cerr)
	// 		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error closing response body: "+cerr.Error())
	// 	}
	// }()

	// // Read the HTML-encoded JSON string from the response body
	// htmlEncodedJSON, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Printf("[GATEWAY] Error reading response body: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
	// 	return
	// }

	// // Decode HTML-encoded JSON string to ResponseData
	// var decodedResponse models.ResponseData
	// if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
	// 	log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
	// 	return
	// }

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] CreatePatient: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusConflict:
		log.Printf("[GATEWAY] CreatePatient: Request failed with conflict status %d", status)
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Patient Create Conflict: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] CreatePatient: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetPatients handles fetching all patients.
func (gc *GatewayController) GetPatients(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get all patients.")

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.GET, utils.PATIENT_FETCH_ALL_PATIENTS_ENDPOINT, utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}

	// // Close the response body explicitly after decoding
	// defer func() {
	// 	if cerr := response.Body.Close(); cerr != nil {
	// 		log.Printf("[GATEWAY] Error closing response body: %v", cerr)
	// 		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error closing response body: "+cerr.Error())
	// 	}
	// }()

	// // Read the HTML-encoded JSON string from the response body
	// htmlEncodedJSON, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Printf("[GATEWAY] Error reading response body: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
	// 	return
	// }

	// // Decode HTML-encoded JSON string to ResponseData
	// var decodedResponse models.ResponseData
	// if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
	// 	log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
	// 	return
	// }

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetPatients: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	default:
		log.Printf("[GATEWAY] GetPatients: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetPatientByID handles fetching a patient by ID.
func (gc *GatewayController) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get patient by ID.")

	// Get PatientID from request params
	patientIDString := mux.Vars(r)[utils.GET_PATIENT_ID_PARAMETER]

	// Convert patientIDString to int64
	patientID, err := strconv.ParseInt(patientIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid patient ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, patientID), utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}

	// // Close the response body explicitly after decoding
	// defer func() {
	// 	if cerr := response.Body.Close(); cerr != nil {
	// 		log.Printf("[GATEWAY] Error closing response body: %v", cerr)
	// 		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error closing response body: "+cerr.Error())
	// 	}
	// }()

	// // Read the HTML-encoded JSON string from the response body
	// htmlEncodedJSON, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Printf("[GATEWAY] Error reading response body: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
	// 	return
	// }

	// // Decode HTML-encoded JSON string to ResponseData
	// var decodedResponse models.ResponseData
	// if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
	// 	log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
	// 	return
	// }

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetPatientByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetPatientByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetPatientByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetPatientByEmail handles fetching a patient by email.
func (gc *GatewayController) GetPatientByEmail(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get patient by email.")

	// Get PatientEmail from request params
	patientEmail := mux.Vars(r)[utils.GET_PATIENT_EMAIL_PARAMETER]

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.PATIENT_FETCH_PATIENT_BY_EMAIL_ENDPOINT, patientEmail), utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}

	// // Close the response body explicitly after decoding
	// defer func() {
	// 	if cerr := response.Body.Close(); cerr != nil {
	// 		log.Printf("[GATEWAY] Error closing response body: %v", cerr)
	// 		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Error closing response body: "+cerr.Error())
	// 	}
	// }()

	// // Read the HTML-encoded JSON string from the response body
	// htmlEncodedJSON, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Printf("[GATEWAY] Error reading response body: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
	// 	return
	// }

	// // Decode HTML-encoded JSON string to ResponseData
	// var decodedResponse models.ResponseData
	// if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
	// 	log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
	// 	return
	// }

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetPatientByEmail: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetPatientByEmail: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetPatientByEmail: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetPatientByUserID handles fetching a patient by user ID.
func (gc *GatewayController) GetPatientByUserID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get patient by UserID.")

	// Get UserID from request params
	userIDString := mux.Vars(r)[utils.GET_PATIENT_USER_ID_PARAMETER]

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid user ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_USER_ID_ENDPOINT, userID), utils.PATIENT_PORT, nil)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// // Close the response body explicitly after decoding
	// defer func() {
	// 	if err := response.Body.Close(); err != nil {
	// 		log.Printf("[GATEWAY] Error closing response body: %v", err)
	// 		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
	// 	}
	// }()

	// // Read the HTML-encoded JSON string from the response body
	// htmlEncodedJSON, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Printf("[GATEWAY] Error reading response body: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
	// 	return
	// }

	// // Decode HTML-encoded JSON string to ResponseData
	// var decodedResponse models.ResponseData
	// if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
	// 	log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
	// 	return
	// }

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetPatientByUserID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetPatientByUserID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetPatientByUserID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// UpdatePatientByID handles updating a patient by ID.
func (gc *GatewayController) UpdatePatientByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to update patient by ID.")

	// Get PatientID from request params
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

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodPut, fmt.Sprintf("%s/%d", utils.PATIENT_UPDATE_PATIENT_BY_ID_ENDPOINT, patientID), utils.PATIENT_PORT, patientData)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// // Close the response body explicitly after decoding
	// defer func() {
	// 	if err := response.Body.Close(); err != nil {
	// 		log.Printf("[GATEWAY] Error closing response body: %v", err)
	// 		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
	// 		return
	// 	}
	// }()

	// // Read the HTML-encoded JSON string from the response body
	// htmlEncodedJSON, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Printf("[GATEWAY] Error reading response body: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
	// 	return
	// }

	// // Decode HTML-encoded JSON string to ResponseData
	// var decodedResponse models.ResponseData
	// if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
	// 	log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
	// 	return
	// }

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] UpdatePatientByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] UpdatePatientByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
		return
	case http.StatusConflict:
		log.Printf("[GATEWAY] UpdatePatientByID: Request failed with conflict status %d", status)
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Patient Update Conflict: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] UpdatePatientByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// DeletePatientByID handles deleting a patient by ID.
func (gc *GatewayController) DeletePatientByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to delete patient by ID.")

	// Get PatientID from request params
	patientIDString := mux.Vars(r)[utils.DELETE_PATIENT_ID_PARAMETER]

	// Convert patientIDString to int64
	patientID, err := strconv.ParseInt(patientIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid patient ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid patient ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodDelete, fmt.Sprintf("%s/%d", utils.PATIENT_DELETE_PATIENT_BY_ID_ENDPOINT, patientID), utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// // Close the response body explicitly after decoding
	// defer func() {
	// 	if err := response.Body.Close(); err != nil {
	// 		log.Printf("[GATEWAY] Error closing response body: %v", err)
	// 		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
	// 		return
	// 	}
	// }()

	// // Read the HTML-encoded JSON string from the response body
	// htmlEncodedJSON, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Printf("[GATEWAY] Error reading response body: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to read response body", "Failed to read response body: "+err.Error())
	// 	return
	// }

	// // Decode HTML-encoded JSON string to ResponseData
	// var decodedResponse models.ResponseData
	// if err := utils.DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
	// 	log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to decode HTML-encoded JSON", "Failed to decode HTML-encoded JSON: "+err.Error())
	// 	return
	// }

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] DeletePatientByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] DeletePatientByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Patient not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] DeletePatientByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

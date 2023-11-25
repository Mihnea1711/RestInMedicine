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

// CreateDoctor handles the creation of a new doctor.
func (gc *GatewayController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to create a doctor.")
	// Take credentials data from the context after validation
	doctorRequest := r.Context().Value(utils.DECODED_DOCTOR_DATA).(*models.DoctorData)

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// check IDUser exists
	userResponse, err := gc.IDMClient.GetUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(doctorRequest.IDUser)}})
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
	response, err := gc.redirectRequestBody(ctx, utils.POST, utils.DOCTOR_CREATE_DOCTOR_ENDPOINT, utils.DOCTOR_PORT, doctorRequest)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

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
			// Respond with the response from the other module
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusConflict:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Doctor Conflict: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetDoctors handles the retrieval of all doctors.
func (gc *GatewayController) GetDoctors(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request to another module
	response, err := gc.redirectRequestBody(ctx, utils.GET, utils.DOCTOR_FETCH_ALL_DOCTORS_ENDPOINT, utils.DOCTOR_PORT, nil)
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
			log.Println("[GATEWAY] Doctors fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetDoctorByID handles the retrieval of a doctor by ID.
func (gc *GatewayController) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	doctorIDString := mux.Vars(r)[utils.GET_DOCTOR_BY_ID_PARAMETER]

	// Convert pacientIDString to int64
	doctorID, err := strconv.ParseInt(doctorIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid doctor ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, doctorID), utils.DOCTOR_PORT, nil)
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
			log.Println("[GATEWAY] Doctor fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetDoctorByEmail handles the retrieval of a doctor by email.
func (gc *GatewayController) GetDoctorByEmail(w http.ResponseWriter, r *http.Request) {
	doctorEmail := mux.Vars(r)[utils.GET_DOCTOR_BY_EMAIL_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.DOCTOR_FETCH_DOCTOR_BY_EMAIL_ENDPOINT, doctorEmail), utils.DOCTOR_PORT, nil)
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
		{
			log.Println("[GATEWAY] Doctor fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Patient Conflict: "+decodedResponse.Error)
		return
	}
}

// GetDoctorByUserID handles the retrieval of a doctor by user ID.
func (gc *GatewayController) GetDoctorByUserID(w http.ResponseWriter, r *http.Request) {
	userIDString := mux.Vars(r)[utils.GET_DOCTOR_BY_USER_ID_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid user ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid doctor user ID", err.Error())
		return
	}

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_USER_ID_ENDPOINT, userID), utils.DOCTOR_PORT, nil)
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
			log.Println("[GATEWAY] Doctor fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// UpdateDoctorByID handles the update of a specific doctor by ID.
func (gc *GatewayController) UpdateDoctorByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to update a doctor.")
	doctorIDString := mux.Vars(r)[utils.UPDATE_DOCTOR_BY_ID_PARAMETER]

	// Convert pacientIDString to int64
	doctorID, err := strconv.ParseInt(doctorIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Error closing response body: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid doctor ID", err.Error())
		return
	}

	// Take credentials data from the context after validation
	doctorData := r.Context().Value(utils.DECODED_DOCTOR_DATA).(*models.DoctorData)

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodPut, fmt.Sprintf("%s/%d", utils.DOCTOR_UPDATE_DOCTOR_BY_ID_ENDPOINT, doctorID), utils.DOCTOR_PORT, doctorData)
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
			log.Println("[GATEWAY] Doctor updated successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle not found case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
			return
		}
	case http.StatusConflict:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Doctor Conflict: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// DeleteDoctorByID handles the deletion of a doctor by ID.
func (gc *GatewayController) DeleteDoctorByID(w http.ResponseWriter, r *http.Request) {
	doctorIDString := mux.Vars(r)[utils.DELETE_DOCTOR_BY_ID_PARAMETER]

	// Convert pacientIDString to int64
	doctorID, err := strconv.ParseInt(doctorIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Error parsing doctor ID string: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Internal Server Error. Invalid doctor ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodDelete, fmt.Sprintf("%s/%d", utils.DOCTOR_DELETE_DOCTOR_BY_ID_ENDPOINT, doctorID), utils.DOCTOR_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error closing response body: %v", err)
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
			log.Println("[GATEWAY] Doctor deleted successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle not found case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

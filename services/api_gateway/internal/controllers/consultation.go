package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// CreateConsultation handles the creation of a new consultation.
func (gc *GatewayController) CreateConsultation(w http.ResponseWriter, r *http.Request) {
	var consultationRequest models.ConsultationData

	// Parse the request body into the ProgramConsultationRequest struct
	if err := json.NewDecoder(r.Body).Decode(&consultationRequest); err != nil {
		// Handle the error (e.g., return a response with an error message)
		log.Printf("[GATEWAY] Invalid appointment request payload: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	responseDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, consultationRequest.IDDoctor), utils.DOCTOR_PORT, nil)
	if errDoctor != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate doctor ID", errDoctor.Error())
		return
	}

	log.Println(responseDoctor)

	defer func() {
		if err := responseDoctor.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing doctor response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		}
	}()

	if responseDoctor.StatusCode != http.StatusOK {
		// Read the HTML-encoded JSON string from the response body
		htmlEncodedJSON, err := io.ReadAll(responseDoctor.Body)
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

		utils.SendErrorResponse(w, responseDoctor.StatusCode, decodedResponse.Message, decodedResponse.Error)
		return
	}

	// Check if appointmentRequest.IDPacient exists
	responsePatient, errPacient := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, consultationRequest.IDPacient), utils.PATIENT_PORT, nil)
	if errPacient != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate patient ID", errPacient.Error())
		return
	}

	defer func() {
		if err := responsePatient.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing patient response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		}
	}()

	if responsePatient.StatusCode != http.StatusOK {
		// Read the HTML-encoded JSON string from the response body
		htmlEncodedJSON, err := io.ReadAll(responsePatient.Body)
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

		utils.SendErrorResponse(w, responsePatient.StatusCode, decodedResponse.Message, decodedResponse.Error)
		return
	}

	// Redirect the request body to programare module to create the programare
	response, err := gc.redirectRequestBody(ctx, utils.POST, utils.CONSULTATION_CREATE_CONSULTATIE_ENDPOINT, utils.CONSULTATION_PORT, consultationRequest)
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
			utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Consultation Conflict: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultations handles the retrieval of all consultations.
func (gc *GatewayController) GetConsultations(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request to another module
	response, err := gc.redirectRequestBody(ctx, utils.GET, utils.CONSULTATION_FETCH_ALL_CONSULTATII_ENDPOINT, utils.CONSULTATION_PORT, nil)
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
			log.Println("[GATEWAY] Consultations fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultationsByDoctorID handles the retrieval of consultations by doctor ID.
func (gc *GatewayController) GetConsultationsByDoctorID(w http.ResponseWriter, r *http.Request) {
	doctorIDString := mux.Vars(r)[utils.GET_CONSULTATION_BY_DOCTOR_ID_PARAMETER]
	doctorID, err := strconv.ParseInt(doctorIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid doctor ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	responseDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, doctorID), utils.DOCTOR_PORT, nil)
	if errDoctor != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate doctor ID", errDoctor.Error())
		return
	}

	defer func() {
		if err := responseDoctor.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing doctor response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		}
	}()

	if responseDoctor.StatusCode != http.StatusOK {
		var responseBody *models.ResponseData
		if err := json.NewDecoder(responseDoctor.Body).Decode(&responseBody); err != nil {
			log.Printf("[GATEWAY] Error decoding doctor response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}

		utils.SendErrorResponse(w, responseDoctor.StatusCode, responseBody.Message, responseBody.Error)
		return
	}

	// Redirect the request body to appointment module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.CONSULTATION_FETCH_CONSULTATIE_BY_DOCTOR_ID_ENDPOINT, doctorID), utils.CONSULTATION_PORT, nil)
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
			log.Println("[GATEWAY] Consultation fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultationsByPacientID handles the retrieval of consultations by pacient ID.
func (gc *GatewayController) GetConsultationsByPacientID(w http.ResponseWriter, r *http.Request) {
	patientIDString := mux.Vars(r)[utils.GET_CONSULTATION_BY_PACIENT_ID_PARAMETER]
	patientID, err := strconv.ParseInt(patientIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid pacient ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDPacient exists
	responsePatient, errPacient := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, patientID), utils.PATIENT_PORT, nil)
	if errPacient != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate patient ID", errPacient.Error())
		return
	}

	defer func() {
		if err := responsePatient.Body.Close(); err != nil {
			log.Printf("[GATEWAY] Error closing patient response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		}
	}()

	if responsePatient.StatusCode != http.StatusOK {
		var responseBody *models.ResponseData
		if err := json.NewDecoder(responsePatient.Body).Decode(&responseBody); err != nil {
			log.Printf("[GATEWAY] Error decoding patient response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}

		utils.SendErrorResponse(w, responsePatient.StatusCode, "Patient validation failed", responseBody.Error)
		return
	}

	// Redirect the request body to appointment module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.CONSULTATION_FETCH_CONSULTATIE_BY_PACIENT_ID_ENDPOINT, patientID), utils.CONSULTATION_PORT, nil)
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
			log.Println("[GATEWAY] Consultation fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultationsByDate handles the retrieval of consultations by date.
func (gc *GatewayController) GetConsultationsByDate(w http.ResponseWriter, r *http.Request) {
	dateStr := mux.Vars(r)[utils.GET_CONSULTATION_BY_DATE_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to appointment module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.CONSULTATION_FETCH_CONSULTATIE_BY_DATE_ENDPOINT, dateStr), utils.CONSULTATION_PORT, nil)
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
			log.Println("[GATEWAY] Consultations fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultationByID handles the retrieval of a consultation by ID.
func (gc *GatewayController) GetConsultationByID(w http.ResponseWriter, r *http.Request) {
	consultationIDString := mux.Vars(r)[utils.GET_CONSULTATION_BY_ID_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.CONSULTATION_FETCH_CONSULTATIE_BY_ID_ENDPOINT, consultationIDString), utils.CONSULTATION_PORT, nil)
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
			log.Println("[GATEWAY] Consultation fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// UpdateConsultationByID handles the update of a specific consultation by ID.
func (gc *GatewayController) UpdateConsultationByID(w http.ResponseWriter, r *http.Request) {
	consultationIDString := mux.Vars(r)[utils.UPDATE_CONSULTATION_BY_ID_PARAMETER]

	var consultationData models.ConsultationData
	if err := json.NewDecoder(r.Body).Decode(&consultationData); err != nil {
		log.Printf("[GATEWAY] Error closing response body: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to programare module to create the consultation
	response, err := gc.redirectRequestBody(ctx, utils.PUT, fmt.Sprintf("%s/%s", utils.CONSULTATION_UPDATE_CONSULTATIE_BY_ID_ENDPOINT, consultationIDString), utils.CONSULTATION_PORT, consultationData)
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
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// DeleteConsultationByID handles the deletion of a consultation by ID.
func (gc *GatewayController) DeleteConsultationByID(w http.ResponseWriter, r *http.Request) {
	consultationIDString := mux.Vars(r)[utils.DELETE_CONSULTATION_BY_ID_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, utils.DELETE, fmt.Sprintf("%s/%s", utils.CONSULTATION_DELETE_CONSULTATIE_BY_ID_ENDPOINT, consultationIDString), utils.CONSULTATION_PORT, nil)
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
			log.Println("[GATEWAY] Consultation deleted successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

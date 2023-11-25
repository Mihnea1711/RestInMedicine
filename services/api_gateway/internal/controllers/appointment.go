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
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// CreateAppointment handles the creation of a new appointment.
func (gc *GatewayController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to create a appointment.")
	// Take appointment data from the context after validation
	appointmentRequest := r.Context().Value(utils.DECODED_APPOINTMENT_DATA).(*models.AppointmentData)

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	responseDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, appointmentRequest.IDDoctor), utils.DOCTOR_PORT, nil)
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
	responsePatient, errPacient := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, appointmentRequest.IDPacient), utils.PATIENT_PORT, nil)
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

	// Redirect the request body to appointment module to create the appointment
	response, err := gc.redirectRequestBody(ctx, utils.POST, utils.APPOINTMENT_CREATE_APPOINTMENT_ENDPOINT, utils.APPOINTMENT_PORT, appointmentRequest)
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
			utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Appointment Conflict: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetAppointments handles the retrieval of all appointments.
func (gc *GatewayController) GetAppointments(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request to another module
	response, err := gc.redirectRequestBody(ctx, utils.GET, utils.APPOINTMENT_FETCH_ALL_APPOINTMENTS_ENDPOINT, utils.APPOINTMENT_PORT, nil)
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
			log.Println("[GATEWAY] Appointments fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetAppointmentByID handles the retrieval of a appointment by ID.
func (gc *GatewayController) GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	appointmentIDString := mux.Vars(r)[utils.GET_APPOINTMENT_ID_PARAMETER]
	appointmentID, err := strconv.ParseInt(appointmentIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid appointment ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.APPOINTMENT_FETCH_APPOINTMENT_BY_ID_ENDPOINT, appointmentID), utils.APPOINTMENT_PORT, nil)
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
			log.Println("[GATEWAY] Appointment fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetAppointmentsByDoctorID handles the retrieval of appointments by doctor ID.
func (gc *GatewayController) GetAppointmentsByDoctorID(w http.ResponseWriter, r *http.Request) {
	doctorIDString := mux.Vars(r)[utils.GET_APPOINTMENT_DOCTOR_ID_PARAMETER]
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

	// Redirect the request body to appointment module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.APPOINTMENT_FETCH_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT, doctorID), utils.APPOINTMENT_PORT, nil)
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
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetAppointmentsByPacientID handles the retrieval of appointments by pacient ID.
func (gc *GatewayController) GetAppointmentsByPacientID(w http.ResponseWriter, r *http.Request) {
	patientIDString := mux.Vars(r)[utils.GET_APPOINTMENT_PATIENT_ID_PARAMETER]
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

	// Redirect the request body to appointment module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.APPOINTMENT_FETCH_APPOINTMENTS_BY_PATIENT_ID_ENDPOINT, patientID), utils.APPOINTMENT_PORT, nil)
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
			log.Println("[GATEWAY] Appointment fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetAppointmentsByDate handles the retrieval of appointments by date.
func (gc *GatewayController) GetAppointmentsByDate(w http.ResponseWriter, r *http.Request) {
	dateStr := mux.Vars(r)[utils.GET_APPOINTMENT_DATE_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to appointment module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.APPOINTMENT_FETCH_APPOINTMENTS_BY_DATE_ENDPOINT, dateStr), utils.APPOINTMENT_PORT, nil)
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
			log.Println("[GATEWAY] Appointments fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetAppointmentsByStatus handles the retrieval of appointments by status.
func (gc *GatewayController) GetAppointmentsByStatus(w http.ResponseWriter, r *http.Request) {
	statusStr := mux.Vars(r)[utils.GET_APPOINTMENT_STATUS_PARAMETER]

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to appointment module
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.APPOINTMENT_FETCH_APPOINTMENTS_BY_STATUS_ENDPOINT, statusStr), utils.APPOINTMENT_PORT, nil)
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
			log.Println("[GATEWAY] Appointments fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// UpdateAppointmentByID handles the update of a specific appointment by ID.
func (gc *GatewayController) UpdateAppointmentByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to update a appointment.")
	appointmentIDString := mux.Vars(r)[utils.GET_APPOINTMENT_ID_PARAMETER]
	appointmentID, err := strconv.ParseInt(appointmentIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid appointment ID", err.Error())
		return
	}

	// Take appointment data from the context after validation
	appointmentData := r.Context().Value(utils.DECODED_APPOINTMENT_DATA).(*models.AppointmentData)

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	responseDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, appointmentData.IDDoctor), utils.DOCTOR_PORT, nil)
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
	responsePatient, errPacient := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, appointmentData.IDPacient), utils.PATIENT_PORT, nil)
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

	// Redirect the request body to appointment module to update the appointment
	response, err := gc.redirectRequestBody(ctx, utils.PUT, fmt.Sprintf("%s/%d", utils.APPOINTMENT_UPDATE_APPOINTMENT_BY_ID_ENDPOINT, appointmentID), utils.APPOINTMENT_PORT, appointmentData)
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
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

// DeleteAppointmentByID handles the deletion of a appointment by ID.
func (gc *GatewayController) DeleteAppointmentByID(w http.ResponseWriter, r *http.Request) {
	appointmentIDString := mux.Vars(r)[utils.DELETE_APPOINTMENT_ID_PARAMETER]
	appointmentID, err := strconv.ParseInt(appointmentIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid appointment ID", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	response, err := gc.redirectRequestBody(ctx, utils.DELETE, fmt.Sprintf("%s/%d", utils.APPOINTMENT_DELETE_APPOINTMENT_BY_ID_ENDPOINT, appointmentID), utils.APPOINTMENT_PORT, nil)
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
			log.Println("[GATEWAY] Appointment deleted successfully")
			utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(response.StatusCode)+". Error: "+decodedResponse.Error)
		return
	}
}

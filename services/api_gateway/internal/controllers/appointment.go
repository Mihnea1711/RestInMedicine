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
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// CreateAppointment handles the creation of a new programare.
func (gc *GatewayController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointmentRequest models.AppointmentData

	// Parse the request body into the MakeAppointmentRequest struct
	if err := json.NewDecoder(r.Body).Decode(&appointmentRequest); err != nil {
		log.Printf("[GATEWAY] Invalid appointment request payload: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	log.Println(appointmentRequest.IDDoctor)

	// Check if appointmentRequest.IDDoctor exists
	responseDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, appointmentRequest.IDDoctor), utils.DOCTOR_PORT, nil)
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
		var responseBody *models.ResponseData
		if err := json.NewDecoder(responseDoctor.Body).Decode(&responseBody); err != nil {
			log.Printf("[GATEWAY] Error decoding doctor response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}

		utils.SendErrorResponse(w, responseDoctor.StatusCode, "Doctor validation failed", responseBody.Error)
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
		var responseBody *models.ResponseData
		if err := json.NewDecoder(responsePatient.Body).Decode(&responseBody); err != nil {
			log.Printf("[GATEWAY] Error decoding patient response body: %v", err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}

		utils.SendErrorResponse(w, responsePatient.StatusCode, "Patient validation failed", responseBody.Error)
		return
	}

	// Redirect the request body to programare module to create the programare
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

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding appointment response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			// Respond with the response from the other module
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusConflict:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, responseBody.Message, responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Something unexpected happened")
		return
	}
}

// GetAppointments handles the retrieval of all programari.
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

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Appointments fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// GetAppointmentByID handles the retrieval of a programare by ID.
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

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Doctor fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, responseBody.Message, responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// GetAppointmentsByDoctorID handles the retrieval of programari by doctor ID.
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

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Doctor fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusConflict, responseBody.Message, responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// GetAppointmentsByPacientID handles the retrieval of programari by pacient ID.
func (gc *GatewayController) GetAppointmentsByPacientID(w http.ResponseWriter, r *http.Request) {
	patientIDString := mux.Vars(r)[utils.GET_APPOINTMENT_PACIENT_ID_PARAMETER]
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
	response, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.APPOINTMENT_FETCH_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT, patientID), utils.APPOINTMENT_PORT, nil)
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
			utils.SendErrorResponse(w, http.StatusConflict, responseBody.Message, responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// GetAppointmentsByDate handles the retrieval of programari by date.
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

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Appointments fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusConflict, responseBody.Message, responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// GetAppointmentsByStatus handles the retrieval of programari by status.
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

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Appointments fetched successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusConflict, responseBody.Message, responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}

// UpdateAppointmentByID handles the update of a specific programare by ID.
func (gc *GatewayController) UpdateAppointmentByID(w http.ResponseWriter, r *http.Request) {
	appointmentIDString := mux.Vars(r)[utils.GET_APPOINTMENT_ID_PARAMETER]
	appointmentID, err := strconv.ParseInt(appointmentIDString, 10, 64)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid appointment ID", err.Error())
		return
	}

	var appointmentData models.AppointmentData
	if err := json.NewDecoder(r.Body).Decode(&appointmentData); err != nil {
		log.Printf("[GATEWAY] Error closing response body: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to programare module to create the programare
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

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			// Respond with the response from the other module
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusConflict:
		{
			// Handle conflict case
			utils.SendErrorResponse(w, http.StatusConflict, responseBody.Message, responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Something unexpected happened")
		return
	}
}

// DeleteAppointmentByID handles the deletion of a programare by ID.
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

	var responseBody *models.ResponseData
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		log.Printf("[GATEWAY] Error decoding response body: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		{
			log.Println("[GATEWAY] Appointment deleted successfully")
			utils.SendMessageResponse(w, http.StatusOK, responseBody.Message, responseBody.Payload)
			return
		}
	case http.StatusNotFound:
		{
			utils.SendErrorResponse(w, http.StatusConflict, responseBody.Message, responseBody.Error)
			return
		}
	default:
		// Handle default case - internal server error
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", errors.New("unexpected status").Error())
		return
	}
}
